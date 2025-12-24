# WeKnora 프로젝트 의존성 패키지 목록

이 문서는 WeKnora 프로젝트에서 사용된 모든 패키지와 라이브러리의 GitHub 주소 및 라이센스 정보를 포함합니다.

## Go 패키지

### 주요 직접 의존성

| 패키지 | GitHub 주소 | 라이센스 |
|--------|------------|----------|
| github.com/PuerkitoBio/goquery | https://github.com/PuerkitoBio/goquery | BSD-3-Clause |
| github.com/chromedp/chromedp | https://github.com/chromedp/chromedp | MIT |
| github.com/elastic/go-elasticsearch/v7 | https://github.com/elastic/go-elasticsearch | Apache-2.0 |
| github.com/elastic/go-elasticsearch/v8 | https://github.com/elastic/go-elasticsearch | Apache-2.0 |
| github.com/gin-contrib/cors | https://github.com/gin-contrib/cors | MIT |
| github.com/gin-gonic/gin | https://github.com/gin-gonic/gin | MIT |
| github.com/go-viper/mapstructure/v2 | https://github.com/go-viper/mapstructure | MIT |
| github.com/golang-jwt/jwt/v5 | https://github.com/golang-jwt/jwt | MIT |
| github.com/golang-migrate/migrate/v4 | https://github.com/golang-migrate/migrate | MIT |
| github.com/google/uuid | https://github.com/google/uuid | BSD-3-Clause |
| github.com/hibiken/asynq | https://github.com/hibiken/asynq | MIT |
| github.com/mark3labs/mcp-go | https://github.com/mark3labs/mcp-go | MIT |
| github.com/minio/minio-go/v7 | https://github.com/minio/minio-go | Apache-2.0 |
| github.com/neo4j/neo4j-go-driver/v6 | https://github.com/neo4j/neo4j-go-driver | Apache-2.0 |
| github.com/ollama/ollama | https://github.com/ollama/ollama | MIT |
| github.com/panjf2000/ants/v2 | https://github.com/panjf2000/ants | MIT |
| github.com/parquet-go/parquet-go | https://github.com/parquet-go/parquet-go | Apache-2.0 |
| github.com/pganalyze/pg_query_go/v6 | https://github.com/pganalyze/pg_query_go | Apache-2.0 |
| github.com/pgvector/pgvector-go | https://github.com/pgvector/pgvector-go | MIT |
| github.com/qdrant/go-client | https://github.com/qdrant/go-client | Apache-2.0 |
| github.com/redis/go-redis/v9 | https://github.com/redis/go-redis | BSD-2-Clause |
| github.com/sashabaranov/go-openai | https://github.com/sashabaranov/go-openai | Apache-2.0 |
| github.com/sirupsen/logrus | https://github.com/sirupsen/logrus | MIT |
| github.com/spf13/viper | https://github.com/spf13/viper | MIT |
| github.com/stretchr/testify | https://github.com/stretchr/testify | MIT |
| github.com/swaggo/files | https://github.com/swaggo/files | MIT |
| github.com/swaggo/gin-swagger | https://github.com/swaggo/gin-swagger | MIT |
| github.com/swaggo/swag | https://github.com/swaggo/swag | MIT |
| github.com/tencentyun/cos-go-sdk-v5 | https://github.com/tencentyun/cos-go-sdk-v5 | MIT |
| github.com/yanyiwu/gojieba | https://github.com/yanyiwu/gojieba | MIT |
| go.opentelemetry.io/otel | https://github.com/open-telemetry/opentelemetry-go | Apache-2.0 |
| go.uber.org/dig | https://github.com/uber-go/dig | MIT |
| golang.org/x/crypto | https://github.com/golang/crypto | BSD-3-Clause |
| golang.org/x/sync | https://github.com/golang/sync | BSD-3-Clause |
| google.golang.org/grpc | https://github.com/grpc/grpc-go | Apache-2.0 |
| google.golang.org/protobuf | https://github.com/protocolbuffers/protobuf-go | BSD-3-Clause |
| gorm.io/driver/postgres | https://github.com/go-gorm/postgres | MIT |
| gorm.io/gorm | https://github.com/go-gorm/gorm | MIT |

### 전체 Go 패키지 목록

프로젝트에서 사용된 전체 Go 패키지는 약 414개입니다. 전체 목록은 `dependencies_list.json` 파일을 참조하세요.

주요 간접 의존성 패키지들:
- github.com/aws/aws-sdk-go (Apache-2.0)
- github.com/docker/docker (Apache-2.0)
- github.com/gin-contrib/sse (MIT)
- github.com/go-playground/validator/v10 (MIT)
- github.com/goccy/go-json (MIT)
- github.com/json-iterator/go (MIT)
- github.com/klauspost/compress (Apache-2.0)
- github.com/modern-go/concurrent (Apache-2.0)
- github.com/pelletier/go-toml/v2 (MIT)
- github.com/pkg/errors (BSD-2-Clause)
- github.com/prometheus/client_golang (Apache-2.0)
- github.com/quic-go/quic-go (MIT)
- github.com/shirou/gopsutil/v4 (BSD-3-Clause)
- github.com/spf13/cobra (Apache-2.0)
- github.com/spf13/pflag (BSD-3-Clause)
- github.com/ugorji/go/codec (MIT)
- go.mongodb.org/mongo-driver (Apache-2.0)

## NPM 패키지 (Frontend)

| 패키지 | NPM 주소 | GitHub 주소 | 라이센스 |
|--------|----------|------------|----------|
| @microsoft/fetch-event-source | https://www.npmjs.com/package/@microsoft/fetch-event-source | https://github.com/microsoft/fetch-event-source | MIT |
| @tsconfig/node22 | https://www.npmjs.com/package/@tsconfig/node22 | https://github.com/tsconfig/bases | MIT |
| @types/dompurify | https://www.npmjs.com/package/@types/dompurify | https://github.com/DefinitelyTyped/DefinitelyTyped | MIT |
| @types/marked | https://www.npmjs.com/package/@types/marked | https://github.com/DefinitelyTyped/DefinitelyTyped | MIT |
| @types/node | https://www.npmjs.com/package/@types/node | https://github.com/DefinitelyTyped/DefinitelyTyped | MIT |
| @types/papaparse | https://www.npmjs.com/package/@types/papaparse | https://github.com/DefinitelyTyped/DefinitelyTyped | MIT |
| @vitejs/plugin-vue | https://www.npmjs.com/package/@vitejs/plugin-vue | https://github.com/vitejs/vite-plugin-vue | MIT |
| @vitejs/plugin-vue-jsx | https://www.npmjs.com/package/@vitejs/plugin-vue-jsx | https://github.com/vitejs/vite-plugin-vue-jsx | MIT |
| @vue/tsconfig | https://www.npmjs.com/package/@vue/tsconfig | https://github.com/vuejs/tsconfig | MIT |
| axios | https://www.npmjs.com/package/axios | https://github.com/axios/axios | MIT |
| dompurify | https://www.npmjs.com/package/dompurify | https://github.com/cure53/DOMPurify | Apache-2.0 |
| highlight.js | https://www.npmjs.com/package/highlight.js | https://github.com/highlightjs/highlight.js | BSD-3-Clause |
| less | https://www.npmjs.com/package/less | https://github.com/less/less.js | Apache-2.0 |
| less-loader | https://www.npmjs.com/package/less-loader | https://github.com/webpack-contrib/less-loader | MIT |
| marked | https://www.npmjs.com/package/marked | https://github.com/markedjs/marked | MIT |
| npm-run-all2 | https://www.npmjs.com/package/npm-run-all2 | https://github.com/mysticatea/npm-run-all | MIT |
| pagefind | https://www.npmjs.com/package/pagefind | https://github.com/CloudCannon/pagefind | MIT |
| papaparse | https://www.npmjs.com/package/papaparse | https://github.com/mholt/PapaParse | MIT |
| pinia | https://www.npmjs.com/package/pinia | https://github.com/vuejs/pinia | MIT |
| swiper | https://www.npmjs.com/package/swiper | https://github.com/nolimits4web/swiper | MIT |
| tdesign-icons-vue-next | https://www.npmjs.com/package/tdesign-icons-vue-next | https://github.com/Tencent/tdesign-icons-vue-next | MIT |
| tdesign-vue-next | https://www.npmjs.com/package/tdesign-vue-next | https://github.com/Tencent/tdesign-vue-next | MIT |
| typescript | https://www.npmjs.com/package/typescript | https://github.com/Microsoft/TypeScript | Apache-2.0 |
| vite | https://www.npmjs.com/package/vite | https://github.com/vitejs/vite | MIT |
| vue | https://www.npmjs.com/package/vue | https://github.com/vuejs/core | MIT |
| vue-i18n | https://www.npmjs.com/package/vue-i18n | https://github.com/intlify/vue-i18n-next | MIT |
| vue-router | https://www.npmjs.com/package/vue-router | https://github.com/vuejs/router | MIT |
| vue-tsc | https://www.npmjs.com/package/vue-tsc | https://github.com/vuejs/language-tools | MIT |
| webpack | https://www.npmjs.com/package/webpack | https://github.com/webpack/webpack | MIT |
| xlsx | https://www.npmjs.com/package/xlsx | https://github.com/SheetJS/sheetjs | Apache-2.0 |

## Python 패키지

### docreader 서브시스템

| 패키지 | PyPI 주소 | GitHub 주소 | 라이센스 |
|--------|-----------|------------|----------|
| antiword | https://pypi.org/project/antiword/ | https://github.com/pyantiword/antiword | GPL-2.0 |
| beautifulsoup4 | https://pypi.org/project/beautifulsoup4/ | https://github.com/waylan/beautifulsoup | MIT |
| cos-python-sdk-v5 | https://pypi.org/project/cos-python-sdk-v5/ | https://github.com/tencentyun/cos-python-sdk-v5 | MIT |
| goose3[all] | https://pypi.org/project/goose3/ | https://github.com/goose3/goose3 | Apache-2.0 |
| grpcio | https://pypi.org/project/grpcio/ | https://github.com/grpc/grpc | Apache-2.0 |
| grpcio-health-checking | https://pypi.org/project/grpcio-health-checking/ | https://github.com/grpc/grpc | Apache-2.0 |
| grpcio-tools | https://pypi.org/project/grpcio-tools/ | https://github.com/grpc/grpc | Apache-2.0 |
| lxml | https://pypi.org/project/lxml/ | https://github.com/lxml/lxml | BSD-3-Clause |
| markdown | https://pypi.org/project/markdown/ | https://github.com/Python-Markdown/markdown | BSD-3-Clause |
| markdownify | https://pypi.org/project/markdownify/ | https://github.com/matthewwithanm/markdownify | MIT |
| markitdown[docx,pdf,xls,xlsx] | https://pypi.org/project/markitdown/ | https://github.com/microsoft/markitdown | MIT |
| minio | https://pypi.org/project/minio/ | https://github.com/minio/minio-py | Apache-2.0 |
| mistletoe | https://pypi.org/project/mistletoe/ | https://github.com/miyuchina/mistletoe | MIT |
| ollama | https://pypi.org/project/ollama/ | https://github.com/ollama/ollama-python | MIT |
| openai | https://pypi.org/project/openai/ | https://github.com/openai/openai-python | MIT |
| paddleocr | https://pypi.org/project/paddleocr/ | https://github.com/PaddlePaddle/PaddleOCR | Apache-2.0 |
| paddlepaddle | https://pypi.org/project/paddlepaddle/ | https://github.com/PaddlePaddle/Paddle | Apache-2.0 |
| pdfplumber | https://pypi.org/project/pdfplumber/ | https://github.com/jsvine/pdfplumber | MIT |
| pillow | https://pypi.org/project/pillow/ | https://github.com/python-pillow/Pillow | HPND |
| playwright | https://pypi.org/project/playwright/ | https://github.com/microsoft/playwright-python | Apache-2.0 |
| protobuf | https://pypi.org/project/protobuf/ | https://github.com/protocolbuffers/protobuf | BSD-3-Clause |
| pydantic | https://pypi.org/project/pydantic/ | https://github.com/pydantic/pydantic | MIT |
| pypdf | https://pypi.org/project/pypdf/ | https://github.com/py-pdf/pypdf | BSD-3-Clause |
| pypdf2 | https://pypi.org/project/pypdf2/ | https://github.com/py-pdf/pypdf2 | BSD-3-Clause |
| python-docx | https://pypi.org/project/python-docx/ | https://github.com/python-openxml/python-docx | MIT |
| requests | https://pypi.org/project/requests/ | https://github.com/psf/requests | Apache-2.0 |
| trafilatura | https://pypi.org/project/trafilatura/ | https://github.com/adbar/trafilatura | GPL-3.0 |
| urllib3 | https://pypi.org/project/urllib3/ | https://github.com/urllib3/urllib3 | MIT |

### mcp-server 서브시스템

| 패키지 | PyPI 주소 | GitHub 주소 | 라이센스 |
|--------|-----------|------------|----------|
| mcp | https://pypi.org/project/mcp/ | https://github.com/modelcontextprotocol/python-sdk | MIT |
| requests | https://pypi.org/project/requests/ | https://github.com/psf/requests | Apache-2.0 |

## 라이센스 요약

### 주요 라이센스 유형

- **MIT License**: 가장 많은 패키지가 사용 (약 60%)
- **Apache-2.0**: 약 25%
- **BSD-3-Clause**: 약 10%
- **GPL-2.0 / GPL-3.0**: 일부 Python 패키지 (antiword, trafilatura)
- **BSD-2-Clause**: 소수 패키지
- **HPND**: Pillow (Python Imaging Library)

### 라이센스 호환성

대부분의 패키지는 MIT 또는 Apache-2.0 라이센스를 사용하여 상용 프로젝트에서 자유롭게 사용 가능합니다. GPL 라이센스를 사용하는 패키지(antiword, trafilatura)는 주의가 필요합니다.

## 참고 사항

1. 이 목록은 프로젝트의 직접 의존성과 주요 간접 의존성을 포함합니다.
2. 전체 간접 의존성 목록은 매우 방대하므로 주요 패키지만 포함했습니다.
3. 라이센스 정보는 각 패키지의 공식 저장소에서 확인한 정보를 기반으로 합니다.
4. 정확한 라이센스 정보는 각 패키지의 공식 저장소의 LICENSE 파일을 참조하세요.
5. 상용 프로젝트에서 사용 시 각 라이센스의 요구사항을 확인하고 준수해야 합니다.

## 업데이트 날짜

이 문서는 2024년에 생성되었으며, 패키지 버전과 라이센스 정보는 프로젝트의 의존성 파일을 기준으로 합니다.

