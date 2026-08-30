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

- [Export OpenTelemetry traces](https://goreleaser.com/customization/telemetry/) of your releases to your own collector (Enterprise plan);
- Run in air-gapped environments with [offline licenses](https://goreleaser.com/pro/#offline-licenses)
  (Business and Enterprise plans);
- [Verify published release assets](https://goreleaser.com/customization/verify/) by re-downloading them and running your own checks;
- Create [macOS installers (`.pkg`)](https://goreleaser.com/customization/package/pkg/);
- Create [Windows installers (`.exe`) with NSIS](https://goreleaser.com/customization/package/nsis/);
- Smart [SemVer tag sorting](https://goreleaser.com/customization/general/git/#semver-sorting);
- Publish to [NPM registries](https://goreleaser.com/customization/publish/npm/);
- [Native sign and notarize](https://goreleaser.com/customization/sign/notarize/#native)
  macOS App Bundles, Disk Images, and Installers;
- Use [AI](https://goreleaser.com/customization/publish/changelog/#enhance-with-ai) to improve/format
  your release notes;
- Further filter artifacts with `if` statements;
- Create [macOS App Bundles (`.app`)](https://goreleaser.com/customization/package/app_bundles/);
- Easily create `alpine`, `apt`, and `yum` repositories with the
  [CloudSmith integration](https://goreleaser.com/customization/publish/cloudsmith/);
- Have [global defaults for homepage, description, etc](https://goreleaser.com/customization/general/metadata/);
- Run [hooks before publishing](https://goreleaser.com/customization/publish/beforepublish/) artifacts;
- Cross publish (e.g. releases to GitLab, pushes Homebrew Tap to GitHub);
- Publish [versioned Homebrew Casks](https://goreleaser.com/customization/publish/homebrew_casks/#versioned-casks)
  and [Formulas](https://goreleaser.com/customization/publish/homebrew_formulas/#versioned-formulas);
- Keep [DockerHub image descriptions up to date](https://goreleaser.com/customization/publish/dockerhub/);
- Create [macOS disk images (`.dmg`)](https://goreleaser.com/customization/package/dmg/);
- Create [Windows installers (`.msi`)](https://goreleaser.com/customization/package/msi/) with msitools or WiX;
- Use `goreleaser release --single-target` to build the whole pipeline for a
  single architecture locally;
- Check boxes in pull request templates;
- [Template entire files](https://goreleaser.com/customization/general/templatefiles/) and add them to the
  release. You can also template files that will be included in archives,
  packages, Docker images, etc...;
- Use the [`.Artifacts`](https://goreleaser.com/customization/general/templates/#artifacts) template
  variable to build more powerful customizations;
- Use extra [template fields](https://goreleaser.com/customization/general/templates/#common-fields-pro)
  and [functions](https://goreleaser.com/customization/general/templates/#functions-pro);
- Upload to multiple [Artifactory](https://goreleaser.com/customization/publish/artifactory/) and
  [HTTP server](https://goreleaser.com/customization/publish/upload/) instances;
- [Split and merge builds](https://goreleaser.com/customization/general/partial/) to speed up your release
  by splitting work, use CGO, or run platform-specific code;
- More [changelog options](https://goreleaser.com/customization/publish/changelog/): Filter commits by path
  & subgroups, group dividers;
- Have custom [before and after hooks for archives](https://goreleaser.com/customization/package/archives/);
- Prepare a release with
  `goreleaser release --prepare`,
  publish and announce it later with
  `goreleaser publish` and
  `goreleaser announce`, or with
  `goreleaser continue`;
- Preview and test your next release's change log with
  `goreleaser changelog`;
- Continuously release [nightly builds](https://goreleaser.com/customization/publish/nightlies/);
- Import pre-built binaries with the
  [`prebuilt` builder](https://goreleaser.com/customization/builds/builders/prebuilt/);
- Rootless build [Docker images](https://goreleaser.com/customization/package/docker/#using-podman)
  and
  [manifests](https://goreleaser.com/customization/package/docker_manifest/#using-podman) with
  [Podman](https://podman.io);
- Easily create `apt`, `yum`, and alpine repositories with the
  [gemfury.io integration](https://goreleaser.com/customization/publish/gemfury/);
- Reuse configuration files with the
  [include keyword](https://goreleaser.com/customization/general/includes/);
- Run commands after the release with
  [global after hooks](https://goreleaser.com/customization/general/hooks/);
- Use GoReleaser within your [monorepo](https://goreleaser.com/customization/monorepo/);
- Create
  [custom template variables](https://goreleaser.com/customization/general/templates/#custom-variables)
  (goes well with [includes](https://goreleaser.com/customization/general/includes/)).

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
