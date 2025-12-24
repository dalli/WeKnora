#!/usr/bin/env python3
"""
패키지의 GitHub 주소와 라이센스 정보를 수집하는 스크립트
"""
import json
import requests
import time
from pathlib import Path

def get_github_url(package_name, package_type):
    """패키지 이름으로부터 GitHub URL 추론"""
    if package_type == 'go':
        if package_name.startswith('github.com/'):
            return f"https://{package_name}"
        elif package_name.startswith('golang.org/x/'):
            return f"https://github.com/golang/{package_name.split('/')[-1]}"
        elif package_name.startswith('google.golang.org/'):
            # google.golang.org 패키지는 보통 다른 저장소에 있음
            return None
        elif package_name.startswith('go.opentelemetry.io/'):
            return f"https://github.com/open-telemetry/{package_name.split('/')[-1]}"
        elif package_name.startswith('gorm.io/'):
            return f"https://github.com/go-gorm/{package_name.split('/')[-1]}"
        elif package_name.startswith('go.uber.org/'):
            return f"https://github.com/uber-go/{package_name.split('/')[-1]}"
    elif package_type == 'npm':
        # npm 패키지는 npmjs.com에서 확인 가능
        return f"https://www.npmjs.com/package/{package_name}"
    elif package_type == 'python':
        # Python 패키지는 PyPI에서 확인 가능
        pkg = package_name.split('[')[0]  # extras 제거
        return f"https://pypi.org/project/{pkg}/"
    return None

def get_license_from_github(gh_url):
    """GitHub API를 통해 라이센스 정보 가져오기"""
    if not gh_url or not gh_url.startswith('https://github.com/'):
        return None
    
    # GitHub API 엔드포인트
    api_url = gh_url.replace('https://github.com/', 'https://api.github.com/repos/')
    license_url = f"{api_url}/license"
    
    try:
        response = requests.get(license_url, timeout=5)
        if response.status_code == 200:
            data = response.json()
            return data.get('license', {}).get('name', 'Unknown')
    except:
        pass
    
    return None

def main():
    deps_file = Path('/Users/dalli/dalli.co.kr/WeKnora/dependencies_list.json')
    if not deps_file.exists():
        print("의존성 파일을 찾을 수 없습니다.")
        return
    
    with open(deps_file, 'r') as f:
        deps = json.load(f)
    
    # 주요 직접 의존성만 처리 (전체는 너무 많음)
    main_go_packages = [
        'github.com/PuerkitoBio/goquery',
        'github.com/chromedp/chromedp',
        'github.com/elastic/go-elasticsearch/v7',
        'github.com/elastic/go-elasticsearch/v8',
        'github.com/gin-contrib/cors',
        'github.com/gin-gonic/gin',
        'github.com/go-viper/mapstructure/v2',
        'github.com/golang-jwt/jwt/v5',
        'github.com/golang-migrate/migrate/v4',
        'github.com/google/uuid',
        'github.com/hibiken/asynq',
        'github.com/mark3labs/mcp-go',
        'github.com/minio/minio-go/v7',
        'github.com/neo4j/neo4j-go-driver/v6',
        'github.com/ollama/ollama',
        'github.com/panjf2000/ants/v2',
        'github.com/parquet-go/parquet-go',
        'github.com/pganalyze/pg_query_go/v6',
        'github.com/pgvector/pgvector-go',
        'github.com/qdrant/go-client',
        'github.com/redis/go-redis/v9',
        'github.com/sashabaranov/go-openai',
        'github.com/sirupsen/logrus',
        'github.com/spf13/viper',
        'github.com/stretchr/testify',
        'github.com/swaggo/files',
        'github.com/swaggo/gin-swagger',
        'github.com/swaggo/swag',
        'github.com/tencentyun/cos-go-sdk-v5',
        'github.com/yanyiwu/gojieba',
        'go.opentelemetry.io/otel',
        'go.uber.org/dig',
        'golang.org/x/crypto',
        'golang.org/x/sync',
        'google.golang.org/grpc',
        'google.golang.org/protobuf',
        'gorm.io/driver/postgres',
        'gorm.io/gorm',
    ]
    
    results = {
        'go': {},
        'npm': {},
        'python': {}
    }
    
    print("Go 패키지 정보 수집 중...")
    for pkg in main_go_packages:
        gh_url = get_github_url(pkg, 'go')
        license = get_license_from_github(gh_url) if gh_url else None
        results['go'][pkg] = {
            'github': gh_url,
            'license': license
        }
        time.sleep(0.1)  # Rate limiting
    
    print("NPM 패키지 정보 수집 중...")
    for pkg in deps['npm']:
        gh_url = get_github_url(pkg, 'npm')
        results['npm'][pkg] = {
            'npm': gh_url,
            'license': None  # npm API로 조회 필요
        }
    
    print("Python 패키지 정보 수집 중...")
    for pkg in deps['python']:
        pkg_clean = pkg.split('[')[0]
        gh_url = get_github_url(pkg_clean, 'python')
        results['python'][pkg] = {
            'pypi': gh_url,
            'license': None  # PyPI API로 조회 필요
        }
    
    output_file = Path('/Users/dalli/dalli.co.kr/WeKnora/package_info.json')
    with open(output_file, 'w') as f:
        json.dump(results, f, indent=2, ensure_ascii=False)
    
    print(f"\n결과가 {output_file}에 저장되었습니다.")

if __name__ == '__main__':
    main()

