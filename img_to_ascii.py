from PIL import Image

# ASCII characters from darkest to brightest
ASCII_CHARS = '█▓▒░ '

def pixel_to_ascii(pixel_brightness):
    index = int(pixel_brightness / 255 * (len(ASCII_CHARS) - 1))
    return ASCII_CHARS[index]

def image_to_ascii(image_path, output_width=80):
    img = Image.open(image_path).convert('L')  # Convert to grayscale

    orig_w, orig_h = img.size
    # Terminal chars are roughly twice as tall as wide, so halve the height
    aspect_ratio = orig_h / orig_w
    output_height = int(output_width * aspect_ratio * 0.45)

    img = img.resize((output_width, output_height))

    ascii_art = []
    pixels = list(img.get_flattened_data() if hasattr(img, 'get_flattened_data') else img.getdata())
    for i in range(0, len(pixels), output_width):
        row = pixels[i:i + output_width]
        ascii_art.append(''.join(pixel_to_ascii(p) for p in row))

    return '\n'.join(ascii_art)

if __name__ == '__main__':
    image_path = r'IMG_6363.JPG'
    output_path = r'ascii_art.txt'

    art = image_to_ascii(image_path, output_width=80)

    # Save to file
    with open(output_path, 'w', encoding='utf-8') as f:
        f.write(art)

    # Print to terminal
    import sys
    sys.stdout.reconfigure(encoding='utf-8')
    print(art)
    print(f'\nASCII art saved to: {output_path}')
