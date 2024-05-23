module.exports = {
    darkMode: "class",
    content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
    theme: {
        extend: {
            colors: {
                dark: {
                    500: '#1f1f1f',
                    800: '#181818',
                    900: '#0d0d0d',
                },
                highlight: {
                    100: '#f7ce6e',
                    500: '#fb923c',
                    900: '#e18103'
                }
            }
        },
    }
};
