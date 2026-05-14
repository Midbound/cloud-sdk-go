# Changelog

## 0.2.0 (2026-05-14)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/Midbound/cloud-sdk-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** api update ([de5730c](https://github.com/Midbound/cloud-sdk-go/commit/de5730cc77e3e69dced9277380748fee5e1efc20))
* **client:** optimize json encoder for internal types ([d86403d](https://github.com/Midbound/cloud-sdk-go/commit/d86403de92e28eca2a31e7ab770e8d753ca54608))
* **go:** add default http client with timeout ([7418809](https://github.com/Midbound/cloud-sdk-go/commit/7418809371c83d4613d07c9841a6403ac9d168eb))
* support setting headers via env ([b3ad6e3](https://github.com/Midbound/cloud-sdk-go/commit/b3ad6e32a6a19e5ec79b03c24e533dd613628549))


### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([7fe6f9b](https://github.com/Midbound/cloud-sdk-go/commit/7fe6f9b6454380b42880e99756228f344bf71267))


### Chores

* avoid embedding reflect.Type for dead code elimination ([5e89840](https://github.com/Midbound/cloud-sdk-go/commit/5e898402d39bf1f48c86340c620936d56cbf7351))
* **internal:** codegen related update ([2116655](https://github.com/Midbound/cloud-sdk-go/commit/21166554bbde7006be50b5cceb14aab16d3ddafd))
* **internal:** codegen related update ([a5d9d17](https://github.com/Midbound/cloud-sdk-go/commit/a5d9d1722bb6eb38356df35717520bfa7b754846))
* **internal:** codegen related update ([4ecee84](https://github.com/Midbound/cloud-sdk-go/commit/4ecee8483fec02f23d06a9c6543541166cecaa3e))
* **internal:** more robust bootstrap script ([a1c1736](https://github.com/Midbound/cloud-sdk-go/commit/a1c1736299175cdd535fafd4a4863d37c94b3e04))
* redact api-key headers in debug logs ([a4ee800](https://github.com/Midbound/cloud-sdk-go/commit/a4ee800142f4f6301b4e817769e0990874095b0c))

## 0.1.0 (2026-03-28)

Full Changelog: [v0.0.2...v0.1.0](https://github.com/Midbound/cloud-sdk-go/compare/v0.0.2...v0.1.0)

### Features

* **api:** api update ([5eb7ddc](https://github.com/Midbound/cloud-sdk-go/commit/5eb7ddce86ebc4bd8914dfaf81f33206c73b9ba3))
* **internal:** support comma format in multipart form encoding ([c45fb78](https://github.com/Midbound/cloud-sdk-go/commit/c45fb782348bc897574d3ba542dde43ef68b20bd))


### Bug Fixes

* prevent duplicate ? in query params ([2e3bcd6](https://github.com/Midbound/cloud-sdk-go/commit/2e3bcd6d08c19bb852d65c642ac400ed4662992e))


### Chores

* **ci:** skip lint on metadata-only changes ([a0a6e55](https://github.com/Midbound/cloud-sdk-go/commit/a0a6e5564a37e775aeb261f3625d018d5cafde62))
* **ci:** skip uploading artifacts on stainless-internal branches ([06111e2](https://github.com/Midbound/cloud-sdk-go/commit/06111e2648427044a80c9c90b1b302cf42f7703f))
* **ci:** support opting out of skipping builds on metadata-only commits ([9068cab](https://github.com/Midbound/cloud-sdk-go/commit/9068cab4fcfca8c8912637c366faf5c849af73db))
* **client:** fix multipart serialisation of Default() fields ([703dd91](https://github.com/Midbound/cloud-sdk-go/commit/703dd91d9f96f345ee9163466cb4fd368c0120d0))
* **internal:** codegen related update ([1cc65a3](https://github.com/Midbound/cloud-sdk-go/commit/1cc65a3702ec577a1b28ecd99d8c420977029cee))
* **internal:** codegen related update ([2793416](https://github.com/Midbound/cloud-sdk-go/commit/27934167c9d0ae2f672f45bad06f4873e911255c))
* **internal:** codegen related update ([aada294](https://github.com/Midbound/cloud-sdk-go/commit/aada2945cfdc2b2859904bf3091285b9f82de4f3))
* **internal:** minor cleanup ([2227351](https://github.com/Midbound/cloud-sdk-go/commit/2227351684575123e8c41ade249ba2265a04257d))
* **internal:** move custom custom `json` tags to `api` ([d5bae36](https://github.com/Midbound/cloud-sdk-go/commit/d5bae3644cd90e8709bb6184cd97ccb41289f0b7))
* **internal:** support default value struct tag ([3b66ede](https://github.com/Midbound/cloud-sdk-go/commit/3b66ede58bffe62dadfdc4bc48a46d24e4641962))
* **internal:** tweak CI branches ([4f321b0](https://github.com/Midbound/cloud-sdk-go/commit/4f321b022f7c27b50a702fcccc8ae7bfa1dc1510))
* **internal:** update gitignore ([198153b](https://github.com/Midbound/cloud-sdk-go/commit/198153bb66f21b327ee248148f3d47e236192f3b))
* **internal:** use explicit returns ([827e67b](https://github.com/Midbound/cloud-sdk-go/commit/827e67bf2c27c02d2f332e60c0097a9b87e31e3e))
* **internal:** use explicit returns in more places ([d209ea1](https://github.com/Midbound/cloud-sdk-go/commit/d209ea1badaadb66f8fe13057bc3fb5bd9ab2d17))
* remove unnecessary error check for url parsing ([d56fdc3](https://github.com/Midbound/cloud-sdk-go/commit/d56fdc35dac404d73eedc9dd01821203011cf8ba))
* **tests:** update webhook tests ([21e7650](https://github.com/Midbound/cloud-sdk-go/commit/21e76508229183e6923fe14f5cdccab8287a3ebb))
* update docs for api:"required" ([50d79e8](https://github.com/Midbound/cloud-sdk-go/commit/50d79e859c74c1452cd9a796e97f1c62432692bc))

## 0.0.2 (2026-02-22)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/Midbound/cloud-sdk-go/compare/v0.0.1...v0.0.2)

### Chores

* configure new SDK language ([1a52c57](https://github.com/Midbound/cloud-sdk-go/commit/1a52c5780b4bd563bfeeac6bc97bf4ccc545b05b))
* update SDK settings ([e877aea](https://github.com/Midbound/cloud-sdk-go/commit/e877aea01108017635419598529d7250bb1c276b))
