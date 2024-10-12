import { useLayoutEffect } from "react";
import { useFeature1 } from "@/src/layers/features/feature1/hooks/useFeature1.tsx";
import { useUsers } from "@/src/layers/features/users/hooks/useUsers.tsx";

function App() {
  const { doSomething, something, somethingStatus } = useFeature1();
  const { localUser, localUserStatus, userCreate } = useUsers();

  useLayoutEffect(() => {
    // doSomething()
    //   .then()
    //   .finally(() => {
    //     console.log("App doSomething (useFeature1) - finally");
    //   });

    userCreate({
      username: "tomasci",
      password: "test",
      email: "test@test.test",
    })
      .then()
      .catch();
  }, []);

  return (
    <>
      <h1>Hello, World!</h1>
      <div>
        <p>{JSON.stringify(localUserStatus)}</p>
        <pre>{JSON.stringify(localUser, null, 4)}</pre>
      </div>
    </>
  );
}

export default App;
