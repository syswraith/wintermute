import { CreateLink } from "@/components/CreateLink";
import { Dashboard } from "@/components/Dashboard";
import { Auth } from "@/components/Auth";
import { Route, Routes, BrowserRouter, Navigate } from "react-router-dom";

function ProtectedRoute({ children }: { children: React.ReactNode }) {
  const isAuthenticated = !!localStorage.getItem("token");
  return isAuthenticated ? <>{children}</> : <Navigate to="/auth" />;
}

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/create" element={<ProtectedRoute><CreateLink /></ProtectedRoute>} />
          <Route path="/dashboard" element={<ProtectedRoute><Dashboard /></ProtectedRoute>} />
          <Route path="/auth" element={<Auth />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
