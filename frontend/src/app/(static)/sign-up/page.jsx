"use client";
import { useState } from "react";
import Link from "next/link";
import "./Signup.css";

export default function Signup() {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
    confirmPassword: "",
    firstName: "",
    lastName: "",
    age: "",
    nickname: "",
    aboutMe: "",
  });
  const [avatar, setAvatar] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");

  const host = process.env.NEXT_PUBLIC_HOST;

  // Email regex for basic validation
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

  const validate = () => {
    if (formData.firstName.length > 10)
      return "First name must be under 10 characters.";
    if (formData.lastName.length > 10)
      return "Last name must be under 10 characters.";
    if (formData.nickname && formData.nickname.length > 10)
      return "Nickname must be under 10 characters.";
    if (!emailRegex.test(formData.email))
      return "Please enter a valid email address.";
    if (!formData.age) return "Please enter your date of birth.";
    if (new Date(formData.age) > new Date())
      return "Date of birth cannot be in the future.";
    if (formData.aboutMe.length > 130)
      return "About Me must be under 130 characters.";
    return "";
  };

  const handleChange = (e) => {
    let { name, value } = e.target;
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
    let data;
    try {
      const submitData = new FormData();
      const dataToSend = {
        ...formData,
        age: formData.age ? +new Date(formData.age) : "",
      };
      submitData.append("userData", JSON.stringify(dataToSend));
      if (avatar) {
        submitData.append("avatar", avatar);
      }
      const response = await fetch(`${host}/register`, {
        method: "POST",
        body: submitData,
      });
      data = await response.json();
      if (!response.ok) {
        setError(data.error || "Failed to register. Please try again.");
        setIsLoading(false);
        return;
      }
      setFormData({
        email: "",
        password: "",
        confirmPassword: "",
        firstName: "",
        lastName: "",
        age: "",
        nickname: "",
        aboutMe: "",
      });
      setAvatar(null);
    } catch {
      setError(data?.error || "Failed to register. Please try again.");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="signup-container">
      <div className="signup-form-container">
        <h2>Create an Account</h2>
        <p className="signup-subtitle">
          Already have an account?{" "}
          <Link href="/login" className="signup-link">
            Login here
          </Link>
        </p>
        {error && (
          <div className="signup-error">
            <p>{error}</p>
          </div>
        )}
        <form className="signup-form" onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="firstName">First Name</label>
            <input
              id="firstName"
              name="firstName"
              type="text"
              required
              value={formData.firstName}
              onChange={handleChange}
              placeholder="Enter your first name"
            />
          </div>

          <div className="form-group">
            <label htmlFor="lastName">Last Name</label>
            <input
              id="lastName"
              name="lastName"
              type="text"
              required
              value={formData.lastName}
              onChange={handleChange}
              placeholder="Enter your last name"
            />
          </div>

          <div className="form-group">
            <label htmlFor="nickname">Nickname (Optional)</label>
            <input
              id="nickname"
              name="nickname"
              type="text"
              value={formData.nickname}
              onChange={handleChange}
              placeholder="Enter your nickname"
            />
          </div>

          <div className="form-group">
            <label htmlFor="age">Date of Birth</label>
            <input
              id="age"
              name="age"
              type="date"
              required
              value={formData.age}
              onChange={handleChange}
            />
          </div>

          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input
              id="email"
              name="email"
              type="email"
              required
              value={formData.email}
              onChange={handleChange}
              placeholder="Enter your email"
            />
          </div>

          <div className="form-group">
            <label htmlFor="password">Password</label>
            <input
              id="password"
              name="password"
              type="password"
              required
              value={formData.password}
              onChange={handleChange}
              placeholder="Enter your password"
            />
          </div>

          <div className="form-group">
            <label htmlFor="password">Confirm Password</label>
            <input
              id="confirmPassword"
              name="confirmPassword"
              type="password"
              required
              value={formData.confirmPassword}
              onChange={handleChange}
              placeholder="Enter your password"
            />
          </div>

          <div className="form-group">
            <label htmlFor="aboutMe">About Me (Optional)</label>
            <textarea
              id="aboutMe"
              name="aboutMe"
              value={formData.aboutMe}
              onChange={handleChange}
              placeholder="Tell us about yourself"
              rows="4"
            />
          </div>

          <div className="form-group">
            <label htmlFor="avatar">Profile Picture (Optional)</label>
            <input
              id="avatar"
              name="avatar"
              type="file"
              accept="image/*"
              onChange={(e) => setAvatar(e.target.files[0])}
            />
          </div>

          <button type="submit" className="signup-button" disabled={isLoading}>
            {isLoading ? "Creating Account..." : "Sign Up"}
          </button>
        </form>
      </div>
    </div>
  );
}
