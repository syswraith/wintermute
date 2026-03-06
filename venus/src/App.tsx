import { CreateLink,  } from "@/components/CreateLink";
import { Dashboard  } from "@/components/Dashboard";
import { Analytics } from "@/components/Analytics";
import { Login } from "@/components/Login";
import { Register } from "@/components/Register";
import { Route, Routes, BrowserRouter } from "react-router-dom";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path='/create'     element={ <CreateLink />  } />
          <Route path='/dashboard'  element={ <Dashboard />   } />
          <Route path='/analytics'  element={ <Analytics />   } />
          <Route path='/login'      element={ <Login />       } />
          <Route path='/register'   element={ <Register />    } />
        </Routes>
      </BrowserRouter>
    </>
  );

}

export default App;

