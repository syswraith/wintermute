import { CreateLink } from "@/components/CreateLink";
import { Dashboard } from "@/components/Dashboard";
import { Analytics } from "@/components/Analytics";
import { Auth } from "@/components/Auth";
import { Route, Routes, BrowserRouter } from "react-router-dom";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/create" element={<CreateLink />} />
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/analytics" element={<Analytics />} />
          <Route path="/auth" element={<Auth />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
