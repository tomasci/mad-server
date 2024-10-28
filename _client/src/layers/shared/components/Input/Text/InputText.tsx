import { Controller, useFormContext } from "react-hook-form";
import { useCallback, useMemo, useState } from "react";

interface Props {
  name: string;
  isPassword?: true;
  isEmail?: true;
  defaultValue?: string;
  placeholder?: string;
}

const InputText = (props: Props) => {
  // props
  const { name, isPassword, isEmail, defaultValue, placeholder } = props;

  // rhf
  const {
    control,
    formState: { errors },
  } = useFormContext();

  // state
  const [showPassword, setShowPassword] = useState(false);

  const inputType = useMemo(() => {
    if (isPassword) {
      return showPassword ? "text" : "password";
    }

    if (isEmail) {
      return "email";
    }

    return "text";
  }, [isPassword, isEmail, showPassword]);

  // functions
  const toggleShowPassword = useCallback(() => {
    setShowPassword((prev) => !prev);
  }, []);

  return (
    <>
      <Controller
        render={({ field }) => {
          return (
            <input
              type={inputType}
              {...field}
              {...(placeholder && { placeholder: placeholder })}
            />
          );
        }}
        name={name}
        control={control}
        defaultValue={defaultValue ?? ""}
      />

      {isPassword && (
        <button type={"button"} onClick={toggleShowPassword}>
          {showPassword ? (
            <>
              <s>o</s>
            </>
          ) : (
            <>o</>
          )}
        </button>
      )}

      {errors[name] && (
        <div style={{ color: "red" }}>{String(errors[name]?.message)}</div>
      )}
    </>
  );
};

export { InputText };
