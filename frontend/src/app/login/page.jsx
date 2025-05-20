"use client";
import { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import "./Login.css";

export default function Login() {
  const router = useRouter();
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");

  const host = process.env.NEXT_PUBLIC_HOST;
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

  const validate = () => {
    if (!emailRegex.test(formData.email)) {
      return "Please enter a valid email address.";
    }
    if (!formData.password) {
      return "Please enter your password.";
    }
    return "";
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    const validationError = validate();
    if (validationError) {
      setError(validationError);
      return;
    }
    setIsLoading(true);

    try {
      console.log('hgksdfghdfghdjfghdkjhsdkfgdkfghk',formData);
      
      const response = await fetch(`${host}/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
        credentials: "include",
      });

      const data = await response.json();

      if (response.ok && data.success === "ok") {
        // Redirect to home page on successful login
        router.push("/");
      } else {
        setError(data.error || "Invalid email or password");
      }
    } catch {
      setError("Login failed. Please try again.");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="login-container">
      <div className="login-form-container">
        <h2>Login to Your Account</h2>
        <p className="login-subtitle">
          Don't have an account?{" "}
          <Link href="/sign-up" className="login-link">
            Sign up here
          </Link>
        </p>
        {error && <div className="login-error">{error}</div>}
        <form className="login-form" onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input
              type="text"
              id="email"
              name="email"
              value={formData.email}
              onChange={handleChange}
              required
              placeholder="Enter your email"
              disabled={isLoading}
            />
          </div>
          <div className="form-group">
            <label htmlFor="password">Password</label>
            <input
              type="password"
              id="password"
              name="password"
              value={formData.password}
              onChange={handleChange}
              required
              placeholder="Enter your password"
              disabled={isLoading}
            />
          </div>

          <div className="login-options">
            <div className="remember-me">
              <input type="checkbox" id="remember" />
              <label htmlFor="remember">Remember me</label>
            </div>
            <Link href="/forgot-password" className="forgot-password">
              Forgot Password?
            </Link>
          </div>

          <button type="submit" className="login-button" disabled={isLoading}>
            {isLoading ? "Logging in..." : "Login"}
          </button>
        </form>
      </div>
    </div>
  );
}
