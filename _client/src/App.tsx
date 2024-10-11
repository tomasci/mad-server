import { useFeature1 } from "./layers/features/feature1/hooks/useFeature1.tsx";
import { useLayoutEffect } from "react";

function App() {
  const { init, doSomething, something, somethingStatus } = useFeature1();

  useLayoutEffect(() => {
    init();
    doSomething()
      .then()
      .finally(() => {
        console.log("App doSomething (useFeature1) - finally");
      });
  }, []);

  return (
    <>
      <h1>Hello, World!</h1>
    </>
  );
}

export default App;
