# FiascoExtension

## Usage

### Encode

```shell
FiascoExtension -a encode -i video.mp4 -o encoded.fco
```

### Decode

```shell
FiascoExtension -a decode -i encoded.fco -o video.mp4
```

## Parameters

- `-a | --action`
    - One of: `[encode, decode]`. Decides if the program should encode a video file to the FIASCO codec of if it should
      decode an encoded FIASCO file.

- `-i | --input`
    - In case of encoding, the path to the video file that should be encoded. All file formats that are recognized by
      your FFmpeg version are applicable here.
    - In case of decoding, the path to the encoded FIASCO file.

- `-o | --output`
    - In case of encoding, the path where the encoded FIASCO file should be written to.
    - In case of decoding, the path where the decoded file should be written to. All file formats that are recognized by
      your FFmpeg version are applicable here.
      
- `-l | --layout`
    - Overrides the tiling layout used to tile videos to groups of pictures during encoding.
    - Overrides the tiling layout used to untile groups of pictures back into a video file during decoding.
    - Format: `\d+x\d+`, where the first number is the width and the second number is the height of the tiling layout.
    - Default value: '4x1'.

- `-f | --fps`
    - Overrides the target fps of a decoded video file.
    - Default value: 25.

- `--ffmpegArgs`
    - Additional arguments to be passed during the FFmpeg call when encoding of decoding.
    - Appended at the end of the call as-is.
    - See [FFmpeg docomentation](https://ffmpeg.org/ffmpeg.html) for details.

- `--fiascoArgs`
    - Additional arguments to be passed during the FIASCO call when encoding of decoding.
    - Appended at the end of the call as-is.
    - See FIASCO man-pages for details.