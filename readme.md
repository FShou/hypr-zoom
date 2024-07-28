# Hypr-Zoom

Hypr-Zoom is a command-line tool written in Go that smoothly animates the cursor zoom factor using different easing functions. It leverages the `hyprctl` command to adjust the cursor zoom and the `ease` library for interpolation.

## Features

- Smoothly animates cursor zoom using a variety of easing functions.
- Configurable animation duration and steps.
- Automatically retrieves the initial zoom factor.

## Installation

To build the project, you need to have Go installed. You can download and install Go from the [official website](https://golang.org/dl/).

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/hypr-zoom.git
    cd hypr-zoom
    ```

2. Build the project:
    ```sh
    go build -o hypr-zoom
    ```

## Usage

The `hypr-zoom` command has several flags to configure the animation:

- `-duration`: Duration of the animation in milliseconds (default: 500)
- `-steps`: Number of steps in the animation (default: 100)
- `-easing`: Easing function to use for the animation (default: InOutExpo)
- `-easingOut`: Easing function to use for the zoom-out animation (optional)
- `-target`: Target zoom factor (default: 2.0)

### Example

```sh
hypr-zoom -easing=InOutExpo -duration=500 -steps=60 -target=1.2
```

This command will animate the zoom factor to 1.2 using the InOutExpo easing function over 500 milliseconds with 60 steps.

```sh
hypr-zoom -easing=OutBack -easingOut=OutExpo 
```

This command will animate the zoom factor using the Out easing function when zooming-in and OutExpo when zooming-out.


### Supported Easing Functions
The following easing functions are supported:

- Linear
- InQuad, OutQuad, InOutQuad
- InCubic, OutCubic, InOutCubic
- InQuart, OutQuart, InOutQuart
- InQuint, OutQuint, InOutQuint
- InSine, OutSine, InOutSine
- InExpo, OutExpo, InOutExpo
- InCirc, OutCirc, InOutCirc
- InBack, OutBack, InOutBack
- InBounce, OutBounce, InOutBounce
- InSquare, OutSquare, InOutSquare

## Acknowledgements
- [fogleman/ease](https://github.com/fogleman/ease) - Easing functions library used in this project.
- [Hyprland](https://hyprland.org) 
