#!/bin/bash

# settings
PROJECT_NAME="my-app-go"

echo "Starting documentation preparation..."

mkdir -p public
cp -r godoc/* public/

# move main page to the root
cp public/pkg/$PROJECT_NAME/index.html public/index.html

# remove localhost (replacing it with project root)
find public -name "*.html" -exec sed -i "s|http://localhost:6060/pkg/|/$PROJECT_NAME/pkg/$PROJECT_NAME/|g" {} +
find public -name "*.html" -exec sed -i "s|http://localhost:6060/|/$PROJECT_NAME/|g" {} +

# make all links absolute relative to the domain
# path in file system: ../../lib/godoc/style.css
# have to become URL: /my-app-go/lib/godoc/style.css
find public -type f \( -name "*.html" -o -name "*.css" -o -name "*.js" \) -exec sed -i "s|\.\./|/$PROJECT_NAME/|g" {} 
# cleanup possible duplicates
find public -type f \( -name "*.html" -o -name "*.css" -o -name "*.js" \) -exec sed -i "s|//$PROJECT_NAME/|/|g" {} +
# one more cleanup. path must start with /my-app-go/
find public -type f \( -name "*.html" -o -name "*.css" -o -name

# fixing links to pkg files (handlers.html -> pkg/my-app-go/handlers.html)
# it's needed for index.html in the root
sed -i "s|href=\"handlers.html\"|href=\"pkg/$PROJECT_NAME/handlers.html\"|g" public/index.html
sed -i "s|href=\"models.html\"|href=\"pkg/$PROJECT_NAME/models.html\"|g" public/index.html
sed -i "s|href=\"repositories.html\"|href=\"pkg/$PROJECT_NAME/repositories.html\"|g" public/index.html

# remove duplicates like /my-app-go/my-app-go/
sed -i 's|/my-app-go/my-app-go/|/my-app-go/|g' public/index.html

# fix style references
sed -i 's|../../lib/|/my-app-go/lib/|g' public/index.html
sed -i 's|../lib/|/my-app-go/lib/|g' public/index.html

echo "Documentation is ready"