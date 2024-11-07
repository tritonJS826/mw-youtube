[Opening shot]
"If you're using a monorepo with multiple Go microservices and need to share common utility code, here’s a quick guide!" 

[Step-by-step walkthrough]

"First, ensure the shared packages are in a pkg directory. Packages in internal can’t be shared across services."

"Next, in each microservice’s go.mod file, use the replace directive to point to the shared package."
```On screen: replace example.com/pkg/shared => ../path/to/shared```

"Finally, import the package using an alias if needed. This lets each service access the shared logic without duplicating code."
```On screen: import sharedpkg "example.com/pkg/shared"```

[Closing shot]
"And that’s it! You’ve now set up shared packages across your microservices. Like and subscribe for more Go tips!"
