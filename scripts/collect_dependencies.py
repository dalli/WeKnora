#!/usr/bin/env python3
"""
의존성 패키지의 GitHub 주소와 라이센스 정보를 수집하는 스크립트
"""
import json
import subprocess
import re
import sys
from pathlib import Path

def get_go_packages():
    """Go 패키지 목록 가져오기"""
    try:
        result = subprocess.run(
            ['go', 'list', '-m', 'all'],
            capture_output=True,
            text=True,
            cwd='/Users/dalli/dalli.co.kr/WeKnora'
        )
        packages = []
        for line in result.stdout.strip().split('\n'):
            if line and not line.startswith('github.com/Tencent/WeKnora'):
                # 버전 정보 제거
                pkg = line.split(' ')[0]
                packages.append(pkg)
        return packages
    except Exception as e:
        print(f"Go 패키지 수집 오류: {e}", file=sys.stderr)
        return []

def get_npm_packages():
    """NPM 패키지 목록 가져오기"""
    try:
        package_json_path = Path('/Users/dalli/dalli.co.kr/WeKnora/frontend/package.json')
        if package_json_path.exists():
            with open(package_json_path, 'r') as f:
                data = json.load(f)
            packages = []
            # dependencies와 devDependencies 모두 포함
            deps = data.get('dependencies', {})
            dev_deps = data.get('devDependencies', {})
            for pkg, version in {**deps, **dev_deps}.items():
                packages.append(pkg)
            return packages
    except Exception as e:
        print(f"NPM 패키지 수집 오류: {e}", file=sys.stderr)
        return []
    return []

def get_python_packages():
    """Python 패키지 목록 가져오기"""
    packages = []
    
    # docreader/pyproject.toml
    try:
        pyproject_path = Path('/Users/dalli/dalli.co.kr/WeKnora/docreader/pyproject.toml')
        if pyproject_path.exists():
            content = pyproject_path.read_text()
            # dependencies 섹션에서 패키지 추출
            in_deps = False
            for line in content.split('\n'):
                if 'dependencies = [' in line:
                    in_deps = True
                    continue
                if in_deps and line.strip().startswith(']'):
                    break
                if in_deps and '>=' in line:
                    pkg = line.split('>=')[0].strip().strip('"').strip("'")
                    packages.append(pkg)
    except Exception as e:
        print(f"docreader 패키지 수집 오류: {e}", file=sys.stderr)
    
    # mcp-server/pyproject.toml
    try:
        pyproject_path = Path('/Users/dalli/dalli.co.kr/WeKnora/mcp-server/pyproject.toml')
        if pyproject_path.exists():
            content = pyproject_path.read_text()
            in_deps = False
            for line in content.split('\n'):
                if 'dependencies = [' in line:
                    in_deps = True
                    continue
                if in_deps and line.strip().startswith(']'):
                    break
                if in_deps and '>=' in line:
                    pkg = line.split('>=')[0].strip().strip('"').strip("'")
                    packages.append(pkg)
    except Exception as e:
        print(f"mcp-server 패키지 수집 오류: {e}", file=sys.stderr)
    
    # mcp-server/requirements.txt
    try:
        req_path = Path('/Users/dalli/dalli.co.kr/WeKnora/mcp-server/requirements.txt')
        if req_path.exists():
            for line in req_path.read_text().split('\n'):
                line = line.strip()
                if line and not line.startswith('#'):
                    pkg = line.split('>=')[0].split('==')[0].split('<=')[0].strip()
                    packages.append(pkg)
    except Exception as e:
        print(f"requirements.txt 수집 오류: {e}", file=sys.stderr)
    
    return list(set(packages))  # 중복 제거

def main():
    print("의존성 패키지 수집 중...")
    
    go_packages = get_go_packages()
    npm_packages = get_npm_packages()
    python_packages = get_python_packages()
    
    print(f"\nGo 패키지: {len(go_packages)}개")
    print(f"NPM 패키지: {len(npm_packages)}개")
    print(f"Python 패키지: {len(python_packages)}개")
    
    # 결과를 JSON 파일로 저장
    result = {
        'go': sorted(go_packages),
        'npm': sorted(npm_packages),
        'python': sorted(python_packages)
    }
    
    output_path = Path('/Users/dalli/dalli.co.kr/WeKnora/dependencies_list.json')
    with open(output_path, 'w') as f:
        json.dump(result, f, indent=2, ensure_ascii=False)
    
    print(f"\n결과가 {output_path}에 저장되었습니다.")
    
    # 간단한 목록 출력
    print("\n=== Go 패키지 (일부) ===")
    for pkg in sorted(go_packages)[:20]:
        print(f"  - {pkg}")
    if len(go_packages) > 20:
        print(f"  ... 외 {len(go_packages) - 20}개")
    
    print("\n=== NPM 패키지 ===")
    for pkg in sorted(npm_packages):
        print(f"  - {pkg}")
    
    print("\n=== Python 패키지 ===")
    for pkg in sorted(python_packages):
        print(f"  - {pkg}")

if __name__ == '__main__':
    main()

