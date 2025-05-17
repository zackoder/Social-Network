# Profile Page Implementation in Social Network

This document explains how the profile page works in our social network application, focusing on how the backend and frontend components interact to create a functional user profile experience.

## Overview

The profile page displays user information, posts, followers, and following data. It handles different types of profiles (public/private) and implements visibility rules based on the relationship between the profile owner and the visitor.

## Backend Components

### Database Tables
- `users` - Stores user profile information (name, email, privacy settings, etc.)
- `posts` - Stores all user posts with privacy settings
- `followers` - Maps follower relationships between users
- `private_post_viewers` - Tracks which users can view private posts

### Key Files and Their Roles

1. **controllers/getposts.go**
   - Contains the `GetProfilePosts` function that handles profile post requests
   - Implements authentication and authorization logic for viewing profiles
   - Manages different visibility rules based on profile privacy and follower status

2. **controllers/profile.go**
   - Manages the following/unfollowing functionality (`HandleFollow`)
   - Handles privacy setting updates (`UpdatePrivacy`)

3. **models/select.go**
   - Contains database query functions that support profile functionality:
     - `GetProfilePost` - Retrieves a user's own posts
     - `IsPrivateProfile` - Checks if a profile is set to private
     - `IsFollower` - Checks if one user follows another
     - `GetPublicAndAlmostPrivatePosts` - Gets posts visible to non-followers
     - `GetAllowedPosts` - Gets posts visible to followers
     - `GetRegistration` - Retrieves user profile information

## Frontend Components

### Key Files (Next.js)
- `/src/app/(dynamic)/profile/[id]/page.jsx` - The profile page component
- Profile-related components for displaying:
  - User information
  - Posts
  - Followers/following lists
  - Follow/unfollow buttons

## Authentication Flow

1. Frontend sends requests with cookies containing session tokens
2. Backend validates tokens using `models.Get_session()`
3. If valid, the user ID is extracted and used for authorization checks

## Profile Viewing Flow

1. User navigates to a profile page with ID parameter
2. Frontend sends request to backend API with profile ID
3. Backend checks:
   - Is the viewer authenticated? (via session token)
   - Is the viewer looking at their own profile?
   - If not their own profile, is the target profile public or private?
   - If private, is the viewer a follower?

4. Based on these checks, the backend returns:
   - All posts for own profile
   - Public and "almost private" posts for non-private profiles
   - Public, "almost private", and specifically shared private posts for followers of private profiles
   - Error for non-followers trying to view private profiles

## Data Flow for Profile Posts

```
Frontend                                 Backend
┌─────────────┐                         ┌──────────────┐
│ Profile Page│                         │ GetProfilePosts│
└─────┬───────┘                         └──────┬───────┘
      │                                        │
      │ GET /api/posts?id={profileId}          │
      ├───────────────────────────────────────►│
      │                                        │
      │                                        │ Check session token
      │                                        │ Check if viewing own profile
      │                                        │
      │                                        │ If own profile:
      │                                        │ Call GetProfilePost()
      │                                        │
      │                                        │ If other profile:
      │                                        │ Check privacy with IsPrivateProfile()
      │                                        │
      │                                        │ If public profile:
      │                                        │ Call GetPublicAndAlmostPrivatePosts()
      │                                        │
      │                                        │ If private profile:
      │                                        │ Check follow status with IsFollower()
      │                                        │
      │                                        │ If follower:
      │                                        │ Call GetAllowedPosts()
      │                                        │
      │                                        │ If not follower:
      │                                        │ Return error
      │                                        │
      │ JSON response with posts or error      │
      │◄───────────────────────────────────────┤
      │                                        │
      │ Render posts or error message          │
      │                                        │
┌─────▼───────┐                         ┌──────▼───────┐
│ Display     │                         │ Return to    │
│ Profile     │                         │ request pool │
└─────────────┘                         └──────────────┘
```

## Complex Areas

### Post Privacy Rules

The application implements three levels of post privacy:

1. **Public** - Visible to all users
2. **Almost Private** - Visible only to followers
3. **Private** - Visible only to specific users selected by the creator

The backend query logic in `GetPublicAndAlmostPrivatePosts` and `GetAllowedPosts` handles these complex visibility rules by joining multiple tables and using conditional queries.

### SQL Query Complexity

Note the SQL query in `GetAllowedPosts`:

```sql
SELECT DISTINCT p.id, p.post_privacy, p.title, p.content, p.user_id, u.first_name, p.imagePath, p.createdAt
FROM posts p
JOIN users u ON p.user_id = u.id
LEFT JOIN private_post_viewers ppv ON p.id = ppv.post_id
WHERE p.user_id = ?
AND (
    p.post_privacy = 'public'
    OR (p.post_privacy = 'almostPrivate')
    OR (p.post_privacy = 'private' AND ppv.viewer_id = ?)
)
ORDER BY p.createdAt DESC
```

This query elegantly handles all three privacy levels in a single database call by:
1. Joining the posts and users tables
2. Left joining with private_post_viewers to check for private post permissions
3. Using a conditional WHERE clause that includes posts based on their privacy settings

### Follow Request System

For private profiles, the following process requires approval:
1. User A sends a follow request to User B with private profile
2. Request is stored in a pending state
3. User B receives a notification about the request
4. User B can accept or decline the follow request
5. If accepted, a new follower relationship is created

For public profiles, following happens automatically without approval.

## Summary

The profile system demonstrates complex interconnections between:
- Authentication (via sessions/cookies)
- Authorization (based on relationships and privacy settings)
- Complex database queries
- Different frontend views based on relationship status

The system ensures users have control over their content privacy while providing a seamless experience for profile viewing.