#!/usr/bin/env python3
import re
import sys


def extract_squash_key(line):
    m = re.match(r'- Update pre-commit hook (\S+) to (v[\d.]+)', line)
    if m:
        return (f'pre-commit:{m.group(1)}', m.group(2))

    m = re.match(r'- Update (\S+(?:\s+\S+)*?) digest to ([a-f0-9]+)', line)
    if m:
        return (f'digest:{m.group(1)}', m.group(2))

    m = re.match(r'- Update go modules', line)
    if m:
        return ('go-modules', '')

    m = re.match(r'- Update github actions?', line)
    if m:
        return ('github-actions', '')

    m = re.match(r'- Update docker images?', line)
    if m:
        return ('docker-images', '')

    m = re.match(r'- Update dependency (\S+) to (v[\d.]+)', line)
    if m:
        return (f'dep:{m.group(1)}', m.group(2))

    m = re.match(r'- Update module (\S+) to (v[\d.]+)', line)
    if m:
        return (f'module:{m.group(1)}', m.group(2))

    m = re.match(r'- Update (actions/\S+) action to (v[\d.]+)', line)
    if m:
        return (f'action:{m.group(1)}', m.group(2))

    m = re.match(r'- Update (\S+) Docker tag to (v[\d.]+)', line)
    if m:
        return (f'docker-tag:{m.group(1)}', m.group(2))

    return None


def format_squashed(key, first_ver, last_ver, count):
    if key.startswith('pre-commit:'):
        name = key.split(':', 1)[1]
        if first_ver and last_ver and first_ver != last_ver:
            return f'- Update pre-commit hook {name} ({first_ver} \u2192 {last_ver})'
        return f'- Update pre-commit hook {name} ({count} updates)'

    if key.startswith('digest:'):
        name = key.split(':', 1)[1]
        return f'- Update {name} digest ({count} updates)'

    if key == 'go-modules':
        return f'- Update go modules ({count} updates)'

    if key == 'github-actions':
        return f'- Update github actions ({count} updates)'

    if key == 'docker-images':
        return f'- Update docker images ({count} updates)'

    if key.startswith('dep:'):
        name = key.split(':', 1)[1]
        if first_ver and last_ver and first_ver != last_ver:
            return f'- Update dependency {name} ({first_ver} \u2192 {last_ver})'
        return f'- Update dependency {name} ({count} updates)'

    if key.startswith('module:'):
        name = key.split(':', 1)[1]
        if first_ver and last_ver and first_ver != last_ver:
            return f'- Update module {name} ({first_ver} \u2192 {last_ver})'
        return f'- Update module {name} ({count} updates)'

    if key.startswith('action:'):
        name = key.split(':', 1)[1]
        if first_ver and last_ver and first_ver != last_ver:
            return f'- Update {name} action ({first_ver} \u2192 {last_ver})'
        return f'- Update {name} action ({count} updates)'

    if key.startswith('docker-tag:'):
        name = key.split(':', 1)[1]
        if first_ver and last_ver and first_ver != last_ver:
            return f'- Update {name} Docker tag ({first_ver} \u2192 {last_ver})'
        return f'- Update {name} Docker tag ({count} updates)'

    return None


def squash_section(lines):
    groups = {}
    group_order = []
    entry_positions = {}

    for i, line in enumerate(lines):
        key_info = extract_squash_key(line)
        if not key_info:
            continue
        key, ver = key_info
        if key not in groups:
            groups[key] = {'first_pos': i, 'first_ver': ver, 'last_ver': ver, 'count': 1}
            group_order.append(key)
        else:
            groups[key]['last_ver'] = ver
            groups[key]['count'] += 1

    squashed_at = {}
    skip_positions = set()

    for key in group_order:
        g = groups[key]
        if g['count'] < 2:
            continue
        squashed_line = format_squashed(key, g['first_ver'], g['last_ver'], g['count'])
        squashed_at[g['first_pos']] = squashed_line

    for i, line in enumerate(lines):
        key_info = extract_squash_key(line)
        if not key_info:
            continue
        key, ver = key_info
        if groups[key]['count'] < 2:
            continue
        if i != groups[key]['first_pos']:
            skip_positions.add(i)

    result = []
    for i, line in enumerate(lines):
        if i in skip_positions:
            continue
        if i in squashed_at:
            result.append(squashed_at[i])
        else:
            result.append(line)

    cleaned = []
    prev_blank = False
    for line in result:
        is_blank = line.strip() == ''
        if is_blank and prev_blank:
            continue
        cleaned.append(line)
        prev_blank = is_blank

    return cleaned


def process_changelog(content):
    lines = content.split('\n')
    result = []
    section_lines = []
    in_section = False

    for line in lines:
        if re.match(r'^###\s+', line):
            if section_lines:
                result.extend(squash_section(section_lines))
                section_lines = []
            result.append(line)
            in_section = True
        elif re.match(r'^##\s+', line):
            if section_lines:
                result.extend(squash_section(section_lines))
                section_lines = []
            result.append(line)
            in_section = False
        elif in_section and line.startswith('- '):
            section_lines.append(line)
        elif in_section and line.strip() == '':
            section_lines.append(line)
        else:
            if section_lines:
                result.extend(squash_section(section_lines))
                section_lines = []
            result.append(line)
            if in_section and not line.startswith('- ') and line.strip() != '':
                in_section = False

    if section_lines:
        result.extend(squash_section(section_lines))

    return '\n'.join(result)


if __name__ == '__main__':
    filepath = sys.argv[1] if len(sys.argv) > 1 else 'CHANGELOG.md'
    with open(filepath, 'r') as f:
        content = f.read()
    result = process_changelog(content)
    with open(filepath, 'w') as f:
        f.write(result)
    print(f'Squashed changelog written to {filepath}')
