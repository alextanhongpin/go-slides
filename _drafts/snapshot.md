Why snapshot
- state of the system, frozen in git commit


Why so much trouble?
- snapshot of the system
- log out intermediate steps
- since everything is committed to the git history, it is easy to know what changed
- able to know the state of the system without actually running the system
- do it in one line of code


Separate fixtures and assertions
- write "steps" that products the snapshots
- read the snapshot and compare
- external tool can parse snapshot and handle the assertions too
- when working with multiple services, this ensures the service you are integrating is visible and transparent
- one thing I enjoy is separating setups and assertions
- load snapshots from git revision - reproducable tests. What if there is a newer version? load the newer version.
- saves storage and staging environment - run tests once locally and produce the snapshot for testing.



What if the type you want to serialize is not json serializable?
just create another struct ...


In the era of LLM, we need more zero-shots example from data.
Snapshots provides the zero shot example, especially when decorated with metadata.


## Meta tags for snapshot

- name
- description
- version
- type information (struct name)
- date last updated
- author
- key-value
-	revision hash
- sha
