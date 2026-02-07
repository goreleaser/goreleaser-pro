<p align="center">
  <img alt="GoReleaser Logo" src="https://raw.githubusercontent.com/goreleaser/artwork/master/goreleaser-pro-round.png" height="200" />
  <h3 align="center">GoReleaser Pro</h3>
  <p align="center">Release engineering, simplified.</p>
</p>

---

## What is GoReleaser Pro?

GoReleaser Pro is a fork of the OSS version you already use every day, with
extra features.

With GoReleaser Pro you can:

- Create macOS [`.pkg` installers](https://goreleaser.com/customization/pkg);
- Create Windows installers with [NSIS](https://goreleaser.com/customization/nsis);
- Smart [SemVer tag sorting](https://goreleaser.com/customization/git#semver-sorting);
- Publish to [NPM registries](https://goreleaser.com/customization/npm);
- Native macOS App Bundles, DMG, and `.pkg` signing and
  [notarization](https://goreleaser.com/customization/notarize#native);
- Use [AI](https://goreleaser.com/customization/changelog#enhance-with-ai) to
  improve/format your release notes;
- Further filter artifacts with `if` statements;
- Create macOS [App Bundles](https://goreleaser.com/customization/app_bundles);
- Easily create `alpine`, `apt`, and `yum` repositories with the [CloudSmith integration](https://goreleaser.com/customization/cloudsmith);
- Have [global defaults for homepage, description, etc](https://goreleaser.com/customization/metadata);
- Run [hooks before publishing](https://goreleaser.com/customization/beforepublish) artifacts;
- Cross publish (e.g. releases to GitLab, pushes Homebrew Tap to GitHub);
- Keep [DockerHub image descriptions up to date](https://goreleaser.com/customization/dockerhub);
- Create [macOS disk images (DMGs)](https://goreleaser.com/customization/dmg);
- Create [Windows installers](https://goreleaser.com/customization/msi);
- Use `goreleaser release --single-target` to build the whole pipeline for a
  single architecture locally;
- Check boxes in pull request templates;
- [Template entire files](https://goreleaser.com/customization/templatefiles) and add them to the
  release. You can also template files that will be included in archives,
  packages, Docker images, etc...;
- Use the [`.Artifacts`](https://goreleaser.com/customization/templates/#artifacts) template
  variable to build more powerful customizations;
- [Split and merge builds](https://goreleaser.com/customization/partial) to speed up your release
  by splitting work, use CGO, or run platform-specific code;
- More [changelog options](https://goreleaser.com/customization/changelog): Filter commits by path
  & subgroups, group dividers;
- Have custom [before and after hooks for archives](https://goreleaser.com/customization/archive/);
- Prepare a release with
  [`goreleaser release --prepare`](https://goreleaser.com/cmd/goreleaser_release/), publish and
  announce it later with
  [`goreleaser publish`](https://goreleaser.com/cmd/goreleaser_publish/) and
  [`goreleaser announce`](https://goreleaser.com/cmd/goreleaser_announce/), or with
  [`goreleaser continue`](https://goreleaser.com/cmd/goreleaser_continue/);
- Preview and test your next release's change log with
  [`goreleaser changelog`](https://goreleaser.com/cmd/goreleaser_changelog/);
- Continuously release [nightly builds](https://goreleaser.com/customization/nightlies/);
- Import pre-built binaries with the
  [`prebuilt` builder](https://goreleaser.com./customization/builds.md#import-pre-built-binaries);
- Rootless build [Docker images](https://goreleaser.com./customization/docker.md#podman) and
  [manifests](https://goreleaser.com./customization/docker_manifest.md#podman) with
  [Podman](https://goreleaser.comhttps://podman.io);
- Easily create `apt` and `yum` repositories with the
  [fury.io integration](https://goreleaser.com/customization/fury/);
- Reuse configuration files with the
  [include keyword](https://goreleaser.com/customization/includes/);
- Run commands after the release with
  [global after hooks](https://goreleaser.com/customization/hooks/);
- Use GoReleaser within your [monorepo](https://goreleaser.com/customization/monorepo/);
- Create
  [custom template variables](https://goreleaser.com/customization/templates/#custom-variables)
  (goes well with [includes](https://goreleaser.com/customization/includes/)).

And more features will be added soon.

The idea is to make this more sustainable. I have invested a lot of time in
GoReleaser, some people contribute (either with code or money), most people
don't. I'm not a big fan of receiving money without giving something in return,
so this is my try at it.

## What's next for GoReleaser OSS?

I'll continue to maintain and add features as I already do. The only difference
is that features that feel more "enterprisy" will be added to the paid version
instead (unless someone contributes the code for it).

## What's in this repository?

This repository contains mainly 2 things:

1. The GoReleaser Pro binaries, packages, and archives to download;
1. The GoReleaser Pro configuration `struct`s, so you can use them to integrate
   with other tools, or to generate YAML from Go.

## More details

You can find more details and buy it [here](https://goreleaser.com/pro/).

---

**✨✨ Thanks for your support! ✨✨**
