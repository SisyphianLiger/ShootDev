# ShootDev

This is my Capstone Project, it contains a client in Vue.JS and a server in Golang, more information can be found in the README.md Section

# Introduction

This is my first time using `Vue.js` and `Nuxt.js`, most of my work has been doing backend things like setting up REST API's and creating servers that do some compute. But, that will not cut it for my idea. First let me introduce you to the main concept of this app.

Simply put, in boot.dev users chat with one another in discord. But a thought had occurred to me. Boot.dev is trying to "gameify" online education. But with every good Game comes a way to communicate with the community. So! Why not, just like battle.net is to Blizzard, make a boot.net...(will workshop the name perhaps). Where users who are already on the site can interact with one another, either in a general channel, or whisper friends, while doing modules, and learning how to be a better programmer.

Turns out, you need a beautiful forward facing app if you are going to have a chat app. And to be honest, UI and UX are not my forte, so I did what any developer would. Take "inspiration" from boot.dev, and make a app called "shoot.dev", the online sister to boot.dev. This project, is certain to have multiple chunks to it but in each section you should be able to hear my thoughts about design choices, maybe some roadblocks, and be able to view the project, either with a gif, or some pictures along the way.

Thank you for taking a look at this, and when the project is completed, I will most likely deploy the app, and write a bash script such that you can try it for yourself on your local machine.

# Chapter 1: Landing Page and Login

The Landing Page, is quite iconic for boot.dev, it shows you a weary mage waking up the metaphorical mountain that is programming. I decided after looking at the many different sections of the Landing page, to add the following, 1. Hero Banner 2. Why Post 3. Reccomendations 4. Footer

For The Hero Banner, I wanted it to look very similar to Boot.dev with my own little twists. All buttons, the sign in and start the chat button lead to a login page, which is still actually on the same page but with the help of a emitter, is only triggered by the button press a user would do. I have not wired up the backend to verify that the user is valid, that will come next.

The Why Post and Reccomendations are mainly flavor text, but the key point from a technical standpoint was to get used to tailwind and learn how grids and flexbox work. Tailwind has quite good documenation, and I was very easily able to mimic what the creators of boot.dev had done, with my own silly twist.

Lastly, the footer, was also created in the same inspiration as Why Post and Reccomendations, however, I wanted to make it clear to users that this is a parody site and that boot.dev is the place where I learned how to do this. Paying homage when homage is due.

Lastly all buttons, as previously mentioned, lead to a simply login page. That page requires a email and password, and will most likely be impoved upon in unison with my backend, to have features, such as lost a password, or create username, etc. The classics.

This chapter had me dive deep into grid, flex, and props and emitters, which I found to be somewhat tricky at times, but nice with typescript. The main solution for all buttons within the landing page and post page is actually a union type in TS. That changes and updates in accordance to button select which then changes the view for the user with `v-if`.

Anyways, I think its better off to just show you the current state of the project so...without further ado!

![output](https://github.com/user-attachments/assets/263ba249-f3f7-49d9-adc7-b52bed6b6e72)

# TODOS

1. We have set up the DB but we need the following
   a. Set up the DB hooks for login
   b. Set up Query For DB
   c. Set up JwT and Auth for User
   d. Set up Date and Time for User
   e. Test Functions
   f. Confirm on Client Side
   g. Finish all Log in Features Client Side
   h. Confirm on Server Side
