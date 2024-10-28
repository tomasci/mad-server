import { useCallback } from "react";
import { FormProvider, useForm } from "react-hook-form";
import { InputText } from "@/src/layers/shared/components/Input/Text/InputText.tsx";
import { User_Create_Input } from "@/src/layers/features/users/types.ts";
import { UserCreateInputValidation } from "@/src/layers/features/users/validation/UserCreateInputValidation.ts";
import { useDebug } from "@/src/layers/features/debug/hooks/useDebug.tsx";
import { zodResolver } from "@hookform/resolvers/zod";
import { useUserCreate } from "@/src/layers/features/users/hooks/useUserCreate.tsx";
import { useUser } from "@/src/layers/features/users/hooks/useUser.tsx";
import "@/src/layers/network/endpoints.ts";
import { SnackbarProvider } from "notistack";

function App() {
  // props
  const debug = useDebug();
  const { localUser, localUserStatus, localUserNetworkError } = useUser();
  const { userCreate } = useUserCreate();

  // rhf
  const formMethods = useForm<User_Create_Input>({
    resolver: zodResolver(UserCreateInputValidation),
  });

  // functions
  const onFormSubmit = useCallback(
    (data: User_Create_Input) => {
      debug.log("onFormSubmit", data);
      userCreate(data).then();
    },
    [debug, userCreate],
  );

  return (
    <>
      <SnackbarProvider />

      <h1>Hello, World!</h1>

      <div>
        <pre>{JSON.stringify(localUser, null, 4)}</pre>
        <pre>{JSON.stringify(localUserStatus, null, 4)}</pre>
      </div>

      <div>
        <FormProvider {...formMethods}>
          <form onSubmit={formMethods.handleSubmit(onFormSubmit)}>
            <div>
              <label>Username:</label>
              <InputText
                name={"username"}
                placeholder={"enter username here"}
              />
            </div>

            <div>
              <label>Email:</label>
              <InputText name={"email"} placeholder={"enter email here"} />
            </div>

            <div>
              <label>Password:</label>
              <InputText
                name={"password"}
                isPassword={true}
                placeholder={"enter password here"}
              />
            </div>

            <div>
              <button type={"submit"}>
                {localUserStatus === "idle" ? "Create account" : "Loading..."}
              </button>
            </div>

            {localUserNetworkError && <div style={{color: "red"}}>
              {localUserNetworkError}
            </div>}
          </form>
        </FormProvider>
      </div>
    </>
  );
}

export default App;
