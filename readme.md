# hypr-zoom

`hypr-zoom` is a command-line tool written in Go that smoothly animates the Hyprland cursor zoom-factor changes using variety of easing functions.
It simply uses the `hyprctl` command to adjust the cursor zoom-factor and the ease library for animation interpolation.

## Why ? 
Hyprland cursor zoom factor changes happen instantly, which can feel ood since Hyprland has cool animations. This little CLI solves that.

## Features

- Smoothly animates cursor zoom-factor using a variety of easing and interpolation functions.
- Configurable animation duration and steps.

## Installation
Grab from release or :
1. Clone the repository:
    ```sh
    git clone https://github.com/FaqihS/hypr-zoom.git
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
- `-interp`: Interpolation Function used for animation(default: Log)

### Example

```sh
hypr-zoom -easing=InOutExpo -duration=500 -steps=60 -target=1.2
```

This command will animate the zoom factor to 1.2 using the InOutExpo easing function over 500 milliseconds with 60 steps.

```sh
hypr-zoom -easing=OutBack -easingOut=OutExpo 
```

This command will animate the zoom factor using the OutBack easing function when zooming-in and OutExpo when zooming-out.

> [!WARNING]  
> Adjust duration and steps wisely.

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

For animation preview [see here](https://github.com/fogleman/ease). 

### Supported Interpolation function
- Linear
- Log (Logarithmic)

## Showcase
### Default
```sh
hypr-zoom
```
https://github.com/user-attachments/assets/261aff62-955f-49e6-9a7b-c5f714389dc4
### InOutCubic
```sh
hypr-zoom -easing=InOutCubic
```
https://github.com/user-attachments/assets/622d4a4b-b805-495d-8178-ad5e130279c2
### OutBack-Inback
```sh
hypr-zoom -duration=600 -steps=150 -easing=OutBack -easingOut=InBack
```
https://github.com/user-attachments/assets/de4f2924-b5e8-43a8-9bbf-6f1ac795687c




## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes. Although i'm not real Go Dev :p

## Acknowledgements
- [fogleman/ease](https://github.com/fogleman/ease) - Easing functions library used in this project.
- [Hyprland](https://hyprland.org) 
