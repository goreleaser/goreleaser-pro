<p align="center">
  <img alt="GoReleaser Logo" src="https://raw.githubusercontent.com/goreleaser/artwork/master/goreleaser-pro-round.png" height="200" />
  <h3 align="center">GoReleaser Pro</h3>
  <p align="center">Deliver Go binaries as fast, easily and pro as possible.</p>
</p>

---

## What is GoReleaser Pro?

GoReleaser Pro is a fork of the OSS version you already use every day, with
extra features.

With GoReleaser Pro you can:

- [Split and merge builds](https://goreleaser.com/customization/partial) to speed up your release
  by splitting work, use CGO, or run platform-specific code;
- [Filter commits by path in the changelog](https://goreleaser.com/customization/changelog);
- Have custom [before and after hooks for achives](https://goreleaser.com/customization/archive/);
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
  [`prebuilt` builder](https://goreleaser.com/customization/build/#import-pre-built-binaries);
- Rootless build [Docker images](https://goreleaser.com/customization/docker/#podman) and
  [manifests](https://goreleaser.com/customization/docker_manifest/#podman) with
  [Podman](https://podman.io);
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

## More details

You can find more details and buy it [here](https://goreleaser.com/pro/).

---

**✨✨ Thanks for your support! ✨✨**
