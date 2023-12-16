# `allalt`

`allalt`, a.k.a, "**all** images deserve an **alt** tag", is a simple
CLI tool that transforms images into words. Designed to provide
text-based descriptions of images for visually impaired users, it
leverages the power of
[GPT-4V](https://openai.com/research/gpt-4v-system-card) to make visual
content accessible in a textual format.

`allalt` serves as a handy utility for web developers and content
creators to generate alt text for images, enhancing web accessibility.

The tool is built to be as user-friendly as possible, with a focus on
ease of use and accessibility.

## Installation

### From source

First install the dependencies:

- Go 1.21 or above.
- make.
- [scdoc](https://git.sr.ht/~sircmpwn/scdoc).

Switch to the latest stable tag, `v0.2.0`, then compile and install:

```bash
git checkout v0.2.0
make
sudo make install
```

## Usage

```bash
$ allalt --help
NAME:
   allalt - describe images for visually impaired users

USAGE:
   allalt [global options] [arguments...]

VERSION:
   0.1.0

GLOBAL OPTIONS:
   --key value, -k value                                    the OpenAI API key to use [$ALLALT_KEY]
   --language value, -l value                               the language to use when describing images (default: "en_US") [$ALLALT_LANGUAGE]
   --context value, -c value                                the context around the image to use when describing images [$ALLALT_CONTEXT]
   --keyword value, -K value [ --keyword value, -K value ]  potential keywords relevant to the image
   --help, -h                                               show help
   --version, -v                                            print the version
```

See _allalt(1)_ after installing for more information.

## Contributing

Anyone can help make `allalt` better. Send patches on the [mailing
list](https://lists.sr.ht/~jamesponddotco/allalt-devel) and report bugs
on the [issue tracker](https://todo.sr.ht/~jamesponddotco/allalt).

You must sign-off your work using `git commit --signoff`. Follow the
[Linux kernel developer's certificate of
origin](https://www.kernel.org/doc/html/latest/process/submitting-patches.html#sign-your-work-the-developer-s-certificate-of-origin)
for more details.

All contributions are made under [the GPL-2.0 license](LICENSE.md).

## Resources

The following resources are available:

- [Support and general discussions](https://lists.sr.ht/~jamesponddotco/allalt-discuss).
- [Patches and development related questions](https://lists.sr.ht/~jamesponddotco/allalt-devel).
- [Instructions on how to prepare patches](https://git-send-email.io/).
- [Feature requests and bug reports](https://todo.sr.ht/~jamesponddotco/allalt).

---

Released under the [GPL-2.0 license](LICENSE.md).
