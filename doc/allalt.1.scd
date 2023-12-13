allalt(1)

# NAME

allalt - describe images for visually impaired users

# SYNOPSIS

*allalt* [OPTIONS] FILE

# DESCRIPTION

*allalt* is a CLI tool that transforms images into words. It leverages GPT-4V
to provide text-based descriptions of images, making visual content accessible
in textual format.

# OPTIONS

*-k*, *--key*
	The OpenAI API key to use to generate descriptions.

*-l*, *--language*
	The language to use when describing the image. Defaults to English.

*-K*, *--keyword*
	Potential keyword relevant to the image. Repeat for multiple keywords.

*-h*, *--help*
	Show help message and quit.

*-v*, *--version*
	Show version number and quit.

# ENVIRONMENT

*ALLALT_KEY*
	Specifies the OpenAI API key to use.

*ALLALT_LANGUAGE*
	Specifies the language to use when describing images.

## EXAMPLES

*Example 1. Describe an image assuming the key is in the environment.*
	$ allalt '/path/to/image/file.jpg'

*Example 2. Describe an image in another language.*
	$ allalt -l 'pt_BR' '/path/to/image/file.jpg'

*Example 3. Describe an image with multiple keywords.*
	$ allalt -K 'sunset' -K 'wallpaper' '/path/to/image/file.png'

# AUTHORS

Maintained by James Pond <james@cipher.host>.
