import { ConsoleStyles, UseDebug_Input } from "../types.ts";

const useDebug = (props: UseDebug_Input) => {
  // props
  const { prefix } = props;

  // functions
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const log = (...args: any[]) => {
    // eslint-disable-next-line no-console
    console.log(
      "%c$%s:%c %s",
      [ConsoleStyles.highlightedGreen, ConsoleStyles.white].join(";"),
      prefix,
      ConsoleStyles.unset,
      ...args,
    );
  };

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const error = (...args: any[]) => {
    // eslint-disable-next-line no-console
    console.error(
      "%c$%s:%c %s",
      [ConsoleStyles.highlightedRed, ConsoleStyles.white].join(";"),
      prefix,
      ConsoleStyles.unset,
      ...args,
    );
  };

  return { log, error };
};

export { useDebug };
