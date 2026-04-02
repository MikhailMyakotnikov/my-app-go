#!/bin/bash

# settings
PROJECT_NAME="my-app-go"

echo "Starting documentation preparation..."

mkdir -p public
cp -r godoc/* public/

# move main page to the root
cp public/pkg/my-app-go/index.html public/index.html

# remove localhost (replacing it with project root)
find public -name "*.html" -exec sed -i 's|http://localhost:6060/pkg/|/my-app-go/pkg/my-app-go/|g' {} +
find public -name "*.html" -exec sed -i 's|http://localhost:6060/|/my-app-go/|g' {} +

# make all links absolute relative to the domain
# path in file system: ../../lib/godoc/style.css
# have to become URL: /my-app-go/lib/godoc/style.css
find public -name "*.html" -exec sed -i 's|\.\./|/my-app-go/|g' {} +

# cleanup possible duplicates
find public -name "*.html" -exec sed -i 's|//my-app-go/|/my-app-go/|g' {} +

# fixing links to pkg files (handlers.html -> pkg/my-app-go/handlers.html)
# it's needed for index.html in the root
sed -i 's|href="handlers.html"|href="pkg/my-app-go/handlers.html"|g' public/index.html
sed -i 's|href="models.html"|href="pkg/my-app-go/models.html"|g' public/index.html
sed -i 's|href="repositories.html"|href="pkg/my-app-go/repositories.html"|g' public/index.html

echo "Documentation is ready"