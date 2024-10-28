import { ConsoleStyles, UseDebug_Input } from "../types.ts";

const useDebugClassic = (props?: UseDebug_Input) => {
  // props
  // const { prefix } = props;
  const prefix = props?.prefix ?? "debug"

  // functions
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const log = (...args: any[]) => {
    // const trace = new Error().stack?.split("\n")[2]?.trim().split(" ")[1];
    const stackTrace = new Error().stack?.split("\n")[2]?.trim();
    const traceParts = stackTrace?.split(" ");
    const traceLink = traceParts ? traceParts[traceParts.length - 1] : "";

    // start a collapsed group in the console
    // eslint-disable-next-line no-console
    console.log(
      `%c$%s:%c %s`,
      [ConsoleStyles.highlightedGreen, ConsoleStyles.white].join(";"),
      prefix,
      ConsoleStyles.unset,
      ...args,
      "",
    );

    if (Array.isArray(args) && args.length > 0) {
      if (typeof args[0] === "string" && args[0].includes("trace")) {
        // eslint-disable-next-line no-console
        console.log("^ trace", traceLink);
      }
    }
  };

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const error = (...args: any[]) => {
    // const trace = new Error().stack?.split("\n")[2]?.trim().split(" ")[1];
    const stackTrace = new Error().stack?.split("\n")[2]?.trim();
    const traceParts = stackTrace?.split(" ");
    const traceLink = traceParts ? traceParts[traceParts.length - 1] : "";

    // start a collapsed group in the console
    // eslint-disable-next-line no-console
    console.log(
      `%c$%s:%c %s`,
      [ConsoleStyles.highlightedRed, ConsoleStyles.white].join(";"),
      prefix,
      ConsoleStyles.unset,
      ...args,
      "",
    );

    if (Array.isArray(args) && args.length > 0) {
      if (typeof args[0] === "string" && args[0].includes("trace")) {
        // eslint-disable-next-line no-console
        console.log("^ trace", traceLink);
      }
    }
  };

  return { log, error };
};

const useDebug = (props?: UseDebug_Input) => {
  const functionName = new Error().stack?.split("\n")[2]?.trim().split(" ")[1];
  return useDebugClassic(functionName ? { prefix: functionName } : props ? props : {});
};

export { useDebug };
