# kabaddi-go
A Golang Port for an old Kabaddi game I made using the [Ebiten Framework](https://ebiten.org). And no, not [that Ebiten](https://anilist.co/anime/14073/Ebiten-Kouritsu-Ebisugawa-Koukou-Tenmonbu), pretty sure that one is full on Ecchi.

This project serves as my learning tool to understand how Golang and its intricacies work. Also an exploration of a "weird" game framework that's probably worse to deal with than PyGame. This readme serves as a dev diary of my journey, maybe I should write a tutorial or blog post after this.

For old game please [watch here](https://youtu.be/93ov3ECs9b0).

# Day 0
Installed go and ebiten, created hello word, watched crash course. There's not really anything much going on this day other than it's when I started thinking about the concept itself. Upgrading Kabaddi to be the real sport would prove to be quite impossible if not just not fun to play in a 2D environment. I have decided to simply porting my game with the exact same rulesets. Maybe with some slight change and definitely a web or mobile variant. Though I don't know how to make the control mobile, maybe swipe gestures and/or onscreen keyboard. Would be fun to try out!

# Day 1
I finished the basic movement of the player with a keyboard along with possibly placing traps, points, and the player. Actually tried to do an OOP-like thinggy with interface but I found out that it's not possible to cast a struct into an interface and back into a struct. Python made it look so easy lol. Got a basic structure going by separating each components of the game. Thinking of dividing the game into phases like before, also found out that go is another pointer hell like c. Note to future self. Maybe don't use X O and < for assets, also think of a way that it can stand on its own without the asset folder (Byte decoder or something like that).

Maybe I should consider to be more in depth, but I couldn't be bothered for now.