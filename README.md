[Create Go App][repo_url]

# Create Go App CLI

[![License][repo_license_img]][repo_license_url]

Create a new production-ready project with **backend** (Golang) by running only CLI command.

Focus on **writing your code** and **thinking of the business-logic**! The CLI
will take care of the rest.

## ⚡️ Quick start

First, [download][go_download_url] and install **Go**. Version `1.20` or
higher is required.

Installation is done by using the [`go install`][go_install_url] command:

```bash
go install github.com/comehere127/create-go-app/cmd/cga@latest
```

Or see the repository's [Release page][repo_releases_url], if you want to
download a ready-made `deb`, `rpm`, `apk` or `Arch Linux` package.

Also, GNU/Linux and macOS users available way to install via
[Homebrew][brew_url]:

```bash
# Tap a new formula:
brew tap create-go-app/tap

# Installation:
brew install create-go-app/tap/cga
```

Let's create a new project via **interactive console UI** (or **CUI** for
short) in current folder:

```bash
cga create
```

That's all you need to know to start! 🎉
## ⚙️ Commands & Options

### `create`

CLI command for create a new project with the interactive console UI.

```bash
cga create [OPTION]
```

| Option | Description                                              | Type   | Default | Required? |
|--------|----------------------------------------------------------|--------|---------|-----------|
| `-t`   | Enables to define custom backend templates. | `bool` | `false` | No        |
## 📝 Production-ready project templates

### Backend

- Backend template with Golang built-in [fe-service] package:
  - `fe-service` — Golang service for frontend.

## ⚠️ License

[`Create Go App CLI`][repo_url] is free and open-source software licensed under
the [Apache 2.0 License][repo_license_url]. Official was
created by [Kent Duong][author] and distributed under
[Creative Commons][repo_cc_url] license (CC BY-SA 4.0 International).

<!-- Go -->

[go_download_url]: https://golang.org/dl/
[go_install_url]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies

<!-- Repository -->

[repo_url]: https://github.com/comehere127/create-go-app
[repo_logo_img]: https://github.com/comehere127/create-go-app/assets/11155743/95024afc-5e3b-4d6f-8c9c-5daaa51d080d
[repo_license_url]: https://github.com/comehere127/create-go-app/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[repo_cc_url]: https://creativecommons.org/licenses/by-sa/4.0/
[repo_releases_url]: https://github.com/comehere127/create-go-app/releases

<!-- Author -->
[author]: https://github.com/comehere127

<!-- Readme links -->
[brew_url]: https://brew.sh/

