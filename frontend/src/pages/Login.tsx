const [loading, setLoading] = useState(false);
const [error, setError] = useState("");

const login = async () => {
  try {
    setLoading(true);
    const res = await API.post("/auth/login", { email, password });
    localStorage.setItem("token", res.data.token);
  } catch {
    setError("Invalid credentials");
  } finally {
    setLoading(false);
  }
};

