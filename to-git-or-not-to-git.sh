curl -s https://learn.01founders.co/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"ialali"}}){id}}"}' | cut -d ".." -f4 |cut -d "}" -f1
