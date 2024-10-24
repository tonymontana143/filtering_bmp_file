# filtering_bmp_file
# Bitmap Tool README

## Abstract

This project introduces a tool called **bitmap** that enables users to:

1. Read and write bitmap image files.
2. Apply basic image processing algorithms.
3. Manipulate and transform bitmap images.

## Context

Bitmap files store images as grids of colored pixels. Each pixel's color is defined using a combination of red, green, and blue values (RGB), typically allocating 24 bits (8 bits per color). Bitmap files are divided into two parts: the **Header** (containing metadata like file size, image dimensions, and color depth) and the **Data** (storing the pixel grid).

This tool enables users to explore, transform, and manipulate bitmap images by reading image properties, mirroring, applying filters, rotating, cropping, and more.

## Requirements

- **Code Formatting**: Your code **MUST** follow the `gofumpt` formatting style. Non-compliance results in a zero grade.
- **Successful Compilation**: The program **MUST** compile using:
  ```sh
  go build -o bitmap .
  ```
- **Robustness**: The program **MUST NOT** panic (e.g., nil-pointer dereference, index out of range). Errors must display clear messages and exit with a non-zero status code.

## Usage

The tool provides several functionalities, each accessible via specific commands and flags:

### 1. Header

The `header` command extracts and displays bitmap file header information.

#### Example:
```sh
./bitmap header sample.bmp
```
**Output:**
```
BMP Header:
- FileType: BM
- FileSizeInBytes: 518456
- HeaderSize: 54
DIB Header:
- DibHeaderSize: 40
- WidthInPixels: 480
- HeightInPixels: 360
- PixelSizeInBits: 24
- ImageSizeInBytes: 518402
```

### 2. Mirroring

The `apply --mirror` flag mirrors the image either horizontally or vertically.

#### Usage:
```sh
./bitmap apply --mirror=horizontal sample.bmp output.bmp
```

Supported options:
- **horizontal**, **h**, **horizontally**, **hor**
- **vertical**, **v**, **vertically**, **ver**

### 3. Filters

The `apply --filter` flag applies various image filters. Multiple filters can be specified and will be applied in sequence.

#### Supported filters:
- **blue**: Retains only the blue channel.
- **red**: Retains only the red channel.
- **green**: Retains only the green channel.
- **grayscale**: Converts the image to grayscale.
- **negative**: Applies a negative filter.
- **pixelate**: Pixelates the image (block size: 20 pixels by default).
- **blur**: Applies a blur effect.

#### Example:
```sh
./bitmap apply --filter=grayscale --filter=blur sample.bmp output.bmp
```

### 4. Rotation

The `apply --rotate` flag rotates the image by specified degrees.

#### Supported options:
- **right**, **90**: Rotate 90° clockwise.
- **left**, **-90**: Rotate 90° counterclockwise.
- **180**: Rotate 180°.
- **270**, **-270**: Rotate 270° clockwise or counterclockwise.

#### Example:
```sh
./bitmap apply --rotate=right --rotate=right sample.bmp output.bmp
```

### 5. Cropping

The `apply --crop` flag crops the image based on pixel offsets and dimensions.

#### Usage:
```sh
./bitmap apply --crop=OffsetX-OffsetY-Width-Height sample.bmp output.bmp
```

- **OffsetX** and **OffsetY** define where the crop starts.
- **Width** and **Height** are optional. If omitted, the crop extends to the edge of the image.

#### Example:
```sh
./bitmap apply --crop=20-20-100-100 sample.bmp output.bmp
```

### 6. Combining Options

All options for the `apply` command can be combined and applied in sequence.

#### Example:
```sh
./bitmap apply --mirror=horizontal --rotate=right --filter=negative sample.bmp output.bmp
```

### 7. Help

The `--help` flag provides usage information and has the highest priority. If combined with other options, only the help message will be displayed.

#### Example:
```sh
./bitmap --help
```

## Error Handling

Errors are displayed with clear messages, and the program exits with a non-zero status code. Examples of possible errors:
- File not found
- Invalid bitmap format
- Exceeding crop dimensions

#### Example:
```sh
./bitmap header invalid.txt
Error: invalid.txt is not a bitmap file
$ echo $?
1
```

## Conclusion

This **bitmap** tool provides extensive functionality for working with bitmap images. It allows users to inspect, modify, and transform images using a variety of operations such as mirroring, filtering, rotating, and cropping.
