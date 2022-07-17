## Tweeter
``` https://www.geeksforgeeks.org/design-twitter-a-system-design-interview-question/ ```

### MVP(minimum viarble product) -> Core Feature
1. The user should be able to tweet in just a few second
2. The user should be able to see Tweet timelie
3. Timeline:
   1. user timeline: User see his/her own tweets and tweets user retweet. Tweets which user see when they visit on their profile
   2. home timeline: This will display the tweets from people user follow. (Tweets when you land on twitter.com)
   3. search timeline: When user search some keywords or #tags and they see the tweets related to that particular keywords.
4. The user should be albe to follow another
5. Users should be able to tweet millions of followers within a few seconds (5 seconds) 
    

### NATIVE SOLUTION

To design a big system like Twitter we will firstly talk about the Naive solution. That will help us in moving towards high-level architecture. You can design a solution for the two things: 
1. Data modeling: You can use a relational database like MySQL and you can consider two tables user table (id, username) and a tweet table[id, content, user(primary key of of user table)]. User information will be stored in the user table and whenever a user will tweet a message it will be store in the tweet table. Two relations are also necessary here. One is the user can follow each other, the other is each feed has a user owner. So there will be a one-to-many relationship between user and tweet table.
2. Serve feeds: You need to fetch all the feeds from all the people a user follows and render them in chronological order. 
3. Limitation of Architecture (Point Out Bottleneck)

You will have to do a big select statement in the tweet table to get all the tweets for a specific user whomsoever he/she is following, and that’s also in chronological order. Doing this every time will create a problem because the tweet table will have huge content with lots of tweets. We need to optimize this solution to solve this problem and for that, we will move to the high-level solution for this problem. Before that let’s understand the characteristics of Twitter first. 

4. Characteristics of Twitter (Traffic)

Twitter has 300M daily active users. On average, every second around 6, 000 tweets are tweeted on Twitter. Every second 6, 00, 000 Queries made to get the timelines. Each user has on average 200 followers and some users like some celebrities have millions of followers. This characteristic of twitter clears the following points…