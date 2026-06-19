# Changelog

## 1.0.0 (2026-06-19)


### Features

* **cli:** add radio selection menu ([3abfeeb](https://github.com/JoaoVictorVM/focuzen/commit/3abfeeb7aa14667739c998913e4c61f9c2bfe30d))
* **cli:** localize the tui with go-i18n ([6bad257](https://github.com/JoaoVictorVM/focuzen/commit/6bad25743710ff87928638b307b0643eb1bad4a9))
* **cli:** play radio streams via beep behind an interface ([dfecff6](https://github.com/JoaoVictorVM/focuzen/commit/dfecff6829425988d5f628fa89fe14f5cfe6106c))
* **cli:** scaffold cli module with bubble tea clock ([f2f2111](https://github.com/JoaoVictorVM/focuzen/commit/f2f211126e02641b73ca4c771c3be6c41e44232b))
* **cli:** style the clock with lip gloss ([88c4b60](https://github.com/JoaoVictorVM/focuzen/commit/88c4b6023d132c5dd3dfe10f12da2ee3fe430216))
* **infra:** add terraform for koyeb deployment ([cbb90eb](https://github.com/JoaoVictorVM/focuzen/commit/cbb90ebffe2bc2bb5b5427205dca47f421e28334))
* **server:** add /api/v1/search endpoint backed by the Searcher port ([ca06d9e](https://github.com/JoaoVictorVM/focuzen/commit/ca06d9e20837609babfe2fae50bb5f077d782cf0))
* **server:** add /download redirect to github releases ([69f0395](https://github.com/JoaoVictorVM/focuzen/commit/69f039594f4569e4ed55c4be1e0d661dd1250730))
* **server:** add env config and youtube search client behind interface ([95092d9](https://github.com/JoaoVictorVM/focuzen/commit/95092d9cedc1edaff636f0bd41242cd2b5859dca))
* **server:** add in-memory LRU cache for search to protect quota ([e3b66b4](https://github.com/JoaoVictorVM/focuzen/commit/e3b66b4865265aed8fd2f1cf27a4272702946a6b))
* **server:** add per-IP rate limiting and security headers ([a5c1a02](https://github.com/JoaoVictorVM/focuzen/commit/a5c1a02acbbca25a8b6fd3ec0826b3b1eaa004f5))
* **server:** scaffold backend with chi and health checks ([964b693](https://github.com/JoaoVictorVM/focuzen/commit/964b693aae764efc55caafd7e586c1cab525291b))
* **server:** serve embedded SPA via embed.FS ([1bc4aba](https://github.com/JoaoVictorVM/focuzen/commit/1bc4aba588e2bd634a7362e1a3005d55cddfd992))
* **web:** add central clock and long-form date block ([b0425c1](https://github.com/JoaoVictorVM/focuzen/commit/b0425c1c07692327c3b39337a6c7d67f8360e04e))
* **web:** add fullscreen toggle to the audio controls ([e01d4a3](https://github.com/JoaoVictorVM/focuzen/commit/e01d4a3381bf8ae6e095e930c066b0de6258c1af))
* **web:** add header with logo, search shell and footer ([0bd112e](https://github.com/JoaoVictorVM/focuzen/commit/0bd112eb20b28fa8d5942945660a7ae915ae2456))
* **web:** add hidden youtube iframe player with audio controls ([f4d0303](https://github.com/JoaoVictorVM/focuzen/commit/f4d030342b5674a313db21929379eab908fac4d6))
* **web:** add i18n (pt-BR/en) with language toggle ([31d323b](https://github.com/JoaoVictorVM/focuzen/commit/31d323b0f6b6519f1d66f9d14146cea4e83d1c90))
* **web:** add light/dark theme via CSS variables with toggle ([4173c93](https://github.com/JoaoVictorVM/focuzen/commit/4173c932244f09978aaeb4e5e8c40129990d953c))
* **web:** add local 'next' queue for playback ([914f3b0](https://github.com/JoaoVictorVM/focuzen/commit/914f3b03ebd85794b617563398d70836dc1cfb68))
* **web:** add previous-track control ([8badd96](https://github.com/JoaoVictorVM/focuzen/commit/8badd96a1a7059fe8aea71923f2b2b36e4f9d5db))
* **web:** add repeat-track toggle ([33780e2](https://github.com/JoaoVictorVM/focuzen/commit/33780e284af844d4b8a94ac5965d171642a494db))
* **web:** redesign volume control with mute toggle ([a6e042f](https://github.com/JoaoVictorVM/focuzen/commit/a6e042f15cc02794be8b821d9fe4d85922552db7))
* **web:** scaffold react app with vite, ts strict and tailwind tokens ([14018e6](https://github.com/JoaoVictorVM/focuzen/commit/14018e6a9a8c6195e2b726dc059c18b0df20ca4c))
* **web:** wire search to backend with results dropdown ([a2a34ea](https://github.com/JoaoVictorVM/focuzen/commit/a2a34ea8c87eb8f8b87e9643d34d6bbb01ca0564))


### Bug Fixes

* **cli:** set user agent for audio streams ([0df5c88](https://github.com/JoaoVictorVM/focuzen/commit/0df5c887dfda7ffc2f8615818a09363acbb24960))
* **server:** correct module path to match github repo owner ([7cc71c0](https://github.com/JoaoVictorVM/focuzen/commit/7cc71c09c10b17e57c213924716b894f09a8ed5c))
