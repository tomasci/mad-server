import { useCallback } from "react";

const useNetwork = () => {
  // props

  // state

  // functions
  const query = useCallback(async () => {
    return new Promise((resolve, reject) => {
      let attempt = 0;
      const maxAttempts = 5;

      const tryQuery = () => {
        setTimeout(() => {
          console.log(`Attempt ${attempt + 1}: useNetwork query`);
          attempt++;
          if (attempt < maxAttempts) {
            tryQuery();
          } else {
            resolve({ someResult: "here" });
          }
        }, 1000);
      };

      tryQuery();
    });
  }, []);

  const mutation = useCallback(async () => {
    return new Promise((resolve) => {
      setTimeout(() => {
        console.log("Network mutation performed");
        resolve({});
      }, 1000);
    });
  }, []);

  return {
    query,
    mutation,
  };
};

export { useNetwork };
