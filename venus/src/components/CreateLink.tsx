import { useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Field, FieldGroup, FieldLabel } from "@/components/ui/field";
import { Input } from "@/components/ui/input";

export function CreateLink() {
  const [longURL, setLongURL] = useState<string>("");
  const [shortURL, setShortURL] = useState<string>("");

  return (
    <div className="flex items-center justify-center h-screen">
      <Card className="w-full max-w-2xl">
        <CardHeader>
          <CardTitle>Create Link</CardTitle>
          <CardDescription>Enter required details</CardDescription>
          <CardAction></CardAction>
        </CardHeader>
        <CardContent>
          <form>
            <FieldGroup>
              <div className="grid grid-cols-1 gap-4">
                <Field>
                  <FieldLabel htmlFor="small-form-name">
                    URL to shorten
                  </FieldLabel>
                  <div className="flex gap-2">
                    <Input
                      id="small-form-name"
                      placeholder="Enter URL"
                      required
                      className="flex-1 h-12 text-lg"
                      value={longURL}
                      onChange={(e) => {
                        setLongURL(e.target.value);
                      }}
                    />
                    <Button
                      className="h-12 px-6"
                      type="button"
                      onClick={async () => {
                        try {
                          const text = await navigator.clipboard.readText();
                          setLongURL(text);
                        } catch (err) {
                          console.error("Clipboard read failed:", err);
                        }
                      }}
                    >
                      Paste
                    </Button>
                  </div>
                </Field>
              </div>

              <div className="grid grid-cols-1 gap-4">
                <Field>
                  <FieldLabel htmlFor="small-form-name">
                    Shortened URL
                  </FieldLabel>
                  <div className="flex gap-2">
                    <Input
                      id="small-form-name"
                      className="flex-1 h-12 text-lg"
                      value={shortURL}
                      readOnly
                    />
                    <Button
                      className="h-12 px-6"
                      type="button"
                      onClick={async () => {
                        try {
                          await navigator.clipboard.writeText(shortURL);
                        } catch (err) {
                          console.error(err);
                        }
                      }}
                    >
                      Copy
                    </Button>
                  </div>
                </Field>
              </div>
              <Field orientation="vertical">
                <Button
                  className="h-12 px-6"
                  type="button"
                  onClick={async () => {
                    const url = "http://localhost:31337/create";
                    const res = await fetch(url, {
                      method: "POST",
                      headers: {
                        "Content-Type": "application/json",
                      },
                      body: JSON.stringify({ longURL }),
                    });

                    if (res.status == 200) {
                      const data = await res.json();
                      setShortURL(data.shortURL || "");
                    }
                  }}
                >
                  Submit
                </Button>
              </Field>
            </FieldGroup>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
