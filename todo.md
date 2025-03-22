we will have to create a Facebook-like social network that will contain the following features:

    Followers
    Profile
    Posts
    Groups
    Notification
    Chats

#Frontend

we will have to use a JS framework.

    Next.js
    Vue.js
    Svelte
    Mithril

#Backend

sessions and cookies are required.
Images handling [JPEG,PNG,GIF]
private chats
the application of migrations and the file organization will be tested.

##docker
the project should consist of two Docker images, one for the backend and another for the frontend

Backend Container:

Create a Docker image for the backend of your social network application. This container will run the server-side logic of your application, handle requests from clients, and interact with the database.

Frontend Container:

Create a Docker image for the frontend of your social network application. This container will serve the client-side code, like HTML, CSS, and JavaScript files, to users browsers. It will also communicate with the backend via HTTP requests.

Tips:

Name your frontend Docker image appropriately.
Make sure that the backend container exposes the necessary ports for communication with the frontend and external clients and the frontend container exposes the appropriate port to serve the frontend content to users' browsers.

#Authentication
Email
Password
First Name
Last Name
Date of Birth
Avatar/Image (Optional)
Nickname (Optional)
About Me (Optional)
Note that the Avatar/Image, Nickname and About Me should be present in the form but the user can skip the filling of those fields.

When the user logins, he/she should stay logged in until he/she chooses a logout option that should be available at all times. For this you will have to implement sessions and cookies.

Followers

users should be able to follow and unfollow other users
if a user whants to follow another thie should send a follow request
The recipient user can then choose to "accept" or "decline" the request, if the recipient user has a public profile (as explained in the next section), this request-and-accept process is bypassed and the user who sent the request automatically starts following the user with the pubic profile.

Profile

Every profile should contain :

User information (every information requested in the register form apart from the Password, obviously)
User activity
Every post made by the user
Followers and following users (display the users that are following the owner of the profile and who he/she is following)

There are two types of profiles: a public profile and a private profile. A public profile will display the information specified above to every user on the social network, while the private profile will only display that same information to their followers only.

When the user is in their own profile it should be available an option that allows the user to turn its profile public or private.

Posts

After a user is logged in he/she can create posts and comments on already created posts. While creating a post or a comment, the user can include an image or GIF.

The user must be able to specify the privacy of the post:

public (all users in the social network will be able to see the post)
almost private (only followers of the creator of the post will be able to see the post)
private (only the followers chosen by the creator of the post will be able to see it)

Groups

A user must be able to create a group. The group should have a title and a description given by the creator and he/she can invite other users to join the group.

The invited users need to accept the invitation to be part of the group. They can also invite other people once they are already part of the group. Another way to enter the group is to request to be in it and only the creator of the group would be allowed to accept or refuse the request.

To make a request to enter a group the user must find it first. This will be possible by having a section where you can browse through all groups.

When in a group, a user can create posts and comment the posts already created. These posts and comments will only be displayed to members of the group.

A user belonging to the group can also create an event, making it available for the other group users. An event should have:

    Title
    Description
    Day/Time
    2 Options (at least):
        Going
        Not going

After creating the event every user can choose one of the options for the event.

Chat

Users should be able to send private messages to other users that they are following or being followed, in other words, at least one of the users must be following the other.

When a user sends a message, the recipient will instantly receive it through Websockets if they are following the sender or if the recipient has a public profile.

It should be able for the users to send emojis to each other.

Groups should have a common chat room, so if a user is a member of the group he/she should be able to send and receive messages to this group chat.

Notifications

A user must be able to see the notifications in every page of the project. New notifications are different from new private messages and should be displayed in a different way!

A user should be notified if he/she:

has a private profile and some other user sends him/her a following request
receives a group invitation, so he can refuse or accept the request
is the creator of a group and another user requests to join the group, so he can refuse or accept the request
is member of a group and an event is created

Every other notification created by you that isn't on the list is welcomed too.
