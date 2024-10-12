interface UseDebug_Input {
  prefix: string;
}

enum ConsoleStyles {
  unset = "\\033[0m",
  red = "color: red",
  green = "color: green",
  blue = "color: blue",
  yellow = "color: yellow",
  white = "color: white",
  highlightedYellow = "background: yellow",
  highlightedGreen = "background: green",
  highlightedRed = "background: red",
}

export type { UseDebug_Input };
export { ConsoleStyles };
