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