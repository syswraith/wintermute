import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";

export function Auth({ className, ...props }: React.ComponentProps<"div">) {
  const navigate = useNavigate();
  const [mode, setMode] = useState<"login" | "register">("login");

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const url =
      mode === "login"
        ? "http://localhost:31337/login"
        : "http://localhost:31337/register";

    const body =
      mode === "login"
        ? { username, password }
        : { username, password, confirmPassword };

    const res = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(body),
    });

    if (res.ok) {
      console.log(`${mode} success`);
      const data = await res.json().catch(() => ({}));
      if (data.token) {
        localStorage.setItem("token", data.token);
      }
      navigate("/create");
    } else {
      console.error(`${mode} failed`);
    }
  };

  return (
    <div
      className={cn(
        "flex min-h-screen items-center justify-center p-6",
        className,
      )}
      {...props}
    >
      <Card className="w-full max-w-md">
        <CardHeader>
          <CardTitle>
            {mode === "login" ? "Login to your account" : "Create an account"}
          </CardTitle>

          <CardDescription>
            {mode === "login"
              ? "Enter your credentials to login"
              : "Enter your details to create an account"}
          </CardDescription>
        </CardHeader>

        <CardContent>
          <form onSubmit={handleSubmit}>
            <FieldGroup>
              {/* Username */}
              <Field>
                <FieldLabel htmlFor="username">Username</FieldLabel>
                <Input
                  id="username"
                  type="text"
                  placeholder="admin"
                  required
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                />
              </Field>

              {/* Password */}
              <Field>
                <FieldLabel htmlFor="password">Password</FieldLabel>
                <Input
                  id="password"
                  type="password"
                  required
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </Field>

              {/* Register only */}
              {mode === "register" && (
                <Field>
                  <FieldLabel htmlFor="confirm">Confirm Password</FieldLabel>
                  <Input
                    id="confirm"
                    type="password"
                    required
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                  />
                </Field>
              )}

              {/* Submit */}
              <Field>
                <Button className="w-full" type="submit">
                  {mode === "login" ? "Login" : "Create Account"}
                </Button>
              </Field>

              {/* Switch Mode */}
              <FieldDescription className="text-center">
                {mode === "login" ? (
                  <>
                    Don&apos;t have an account?{" "}
                    <button
                      type="button"
                      className="underline"
                      onClick={() => setMode("register")}
                    >
                      Sign up
                    </button>
                  </>
                ) : (
                  <>
                    Already have an account?{" "}
                    <button
                      type="button"
                      className="underline"
                      onClick={() => setMode("login")}
                    >
                      Login
                    </button>
                  </>
                )}
              </FieldDescription>
            </FieldGroup>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
