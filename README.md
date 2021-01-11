# licengine
serverside licensing engine - the sample of lic engine for software distribution, which aims to provide the sample how it could be realised in most simple way

# improvements
for ideal, you should use any db framework or provide file based user list with any auth data

hardcode sucks and still in the code only for sample

# why should I use/code own lic server against already realized services like auth.gg (not adv)?
well, you won't, but for understanding how it can be designed - I submitted my own template
(optionally: you have mental problems about distribution service security)

you can review it and start an issue, it would help me/other users in future

# how it works
user auth -> retrieving software list -> downloading software from server

every software have an external name and internal

external name fine for widgets like dropdown lists or whatelse, when internal name is using for file, which is stored on the disk (for example)

# what I need for working with it?
well, the sample is providing only API and you need a client

template for client are API requests for user auth and files download

# features
- user auth based on HEAD requests
- ready to use solution
- time based subscriptions
