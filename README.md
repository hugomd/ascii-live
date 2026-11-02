# ascii-live

A project for hosting curl-based animations, all in one place, and a follow up to [`parrot.live`](https://github.com/hugomd/parrot.live).

Any animations you want to add are welcome! 🎉

Try it out in your terminal:
```bash
curl ascii.live/parrot
```

<img src="./demo.gif" width="400"/>

## Upload an image and share live terminal art

PNG and JPEG images can be converted into terminal art and served from a shareable URL. The uploaded art is kept in memory by the running server, so it is available while that server instance is running.

Upload an image with curl:

```bash
curl -X POST -F "image=@photo.png" https://ascii.live/upload
```

The response contains an `url` and a ready-to-copy `command`:

```json
{
  "id": "a1b2c3d4e5f6",
  "format": "png",
  "url": "https://ascii.live/art/a1b2c3d4e5f6",
  "command": "curl https://ascii.live/art/a1b2c3d4e5f6"
}
```

Anyone can run the returned command in a terminal. Uploads are limited to PNG/JPEG files up to 5 MiB and are rendered at a maximum width of 80 terminal columns.

## Running locally
The Go toolchain is managed by [Hermit](https://github.com/cashapp/hermit).
Activate the environment to automatically download and use the pinned Go version:
```bash
source bin/activate-hermit
```

To run the server locally on port `8080`, run:
```bash
go run main.go
```

For a local upload:

```bash
curl -X POST -F "image=@photo.jpg" http://localhost:8080/upload
curl http://localhost:8080/art/<id>
```

## Running in Docker
```bash
docker run -p 8080:8080 hugomd/ascii-live:latest
```

## Adding frames
* [Fork this repository](https://github.com/hugomd/ascii-live/fork)
* Create a new frame file in [`/frames`](./frames), call it the name of your frames/animation, e.g. `parrot.go`
* Inside your new file, create an exported list of frames, e.g.
```Golang
package frames

// This is the value stored in the FrameMap
var MyAnimation = DefaultFrameType(myAnimationFrames)

var myAnimationFrames = []string{
  `Frame1`,
  `Frame2`,
  `Frame3`,
}
```
* In [`./frames/frames.go`](./frames/frames.go), add your animation to the `FrameMap`
* Commit and push your changes, and make a PR! If this is your first time making a PR, [check GitHub's help page on the topic](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request)

## Frame Contributions
Here's a list of lovely people who have contributed frames to this project:

| Contributor                                      | Frame File                                    | Repository                                                      |
|--------------------------------------------------|-----------------------------------------------|-----------------------------------------------------------------|
| [hexrcs](https://github.com/hexrcs)              | [`forrest.go`](./frames/forrest.go)           | [`run-forrest-run`](https://github.com/hexrcs/run-forrest-run)  |
| [jmhobbs](https://github.com/jmhobbs)            | [`parrot.go`](./frames/parrot.go)             | [`terminal-parrot`](https://github.com/jmhobbs/terminal-parrot) |
| [ahs3n](https://github.com/ahs3n)                | [`knot.go`](./frames/knot.go)                 | [`torus-knot`](https://github.com/ahs3n/torus-knot/)            |

## Related Projects
* [parrot.live](https://github.com/hugomd/parrot.live)
* [terminal-parrot](https://github.com/jmhobbs/terminal-parrot)
* [cultofthepartyparrot.com](https://github.com/jmhobbs/cultofthepartyparrot.com)
* [parrotsay](https://github.com/matheuss/parrotsay)
