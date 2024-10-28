interface UseDebug_Input {
  prefix?: string;
}

enum ConsoleStyles {
  unset = "\\033[0m",
  red = "color: red",
  green = "color: green",
  blue = "color: blue",
  yellow = "color: yellow",
  white = "color: white",
  highlightedYellow = "background: yellow",
  highlightedRed = "background: red",
  highlightedGreen = "background: green",
  highlightedBlue = "background: blue",
}

export type { UseDebug_Input };
export { ConsoleStyles };
