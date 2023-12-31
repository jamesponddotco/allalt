package openai_test

import (
	"context"
	"os"
	"testing"

	"git.sr.ht/~jamesponddotco/allalt/internal/openai"
)

func TestClient_Do(t *testing.T) {
	t.Parallel()

	key, found := os.LookupEnv("ALLALT_KEY")
	if !found || key == "" {
		t.Skip("missing required environment variable: ALLALT_KEY")
	}

	tests := []struct {
		name      string
		image     string
		wantError bool
	}{
		{
			name:      "Valid image",
			image:     "data:image/gif;base64,R0lGODlhAQABAPcAAACAAAB/AAB/AAB+AAB+AAB8AAB7AAB6AAB4AAB2AAB0AAFyAQFvAQJtAgJqAgNnAwRkBARgBAVdBQZZBgdVBwlRCQpNCgtJCw1FDQ9BDxE8ERM4ExUzFRcuFxopGh0kHSAgICEhISIiIiMjIyQkJCUlJSYmJicnJygoKCkpKSoqKisrKywsLC0tLS4uLi8vLzAwMDExMTIyMjMzMzQ0NDU1NTY2Njc3Nzg4ODk5OTo6Ojs7Ozw8PD09PT4+Pj8/P0BAQEFBQUJCQkNDQ0REREVFRUZGRkdHR0hISElJSUpKSktLS0xMTE1NTU5OTk9PT1BQUFFRUVJSUlNTU1RUVFVVVVZWVldXV1hYWFlZWVpaWltbW1xcXF1dXV5eXl9fX2BgYGFhYWJiYmNjY2RkZGVlZWZmZmdnZ2hoaGlpaWpqamtra2xsbG1tbW5ubm9vb3BwcHFxcXJycnNzc3R0dHV1dXZ2dnd3d3h4eHl5eXp6ent7e3x8fH19fX5+fn9/f4CAgIGBgYKCgoODg4SEhIWFhYaGhoeHh4iIiImJiYqKiouLi4yMjI2NjY6Ojo+Pj5CQkJGRkZKSkpOTk5SUlJWVlZaWlpeXl5iYmJmZmZqampubm5ycnJ2dnZ6enp+fn6CgoKGhoaKioqOjo6SkpKWlpaampqenp6ioqKmpqaqqqqurq6ysrK2tra6urq+vr7CwsLGxsbKysrOzs7S0tLW1tba2tre3t7i4uLm5ubq6uru7u7y8vL29vb6+vr+/v8DAwMHBwcLCwsPDw8TExMXFxcbGxsfHx8jIyMnJycrKysvLy8zMzM3Nzc7Ozs/Pz9DQ0NHR0dLS0tPT09TU1NXV1dbW1tfX19jY2NnZ2dra2tvb29zc3N3d3d7e3t/f3+Dg4OHh4eLi4uPj4+Tk5OXl5ebm5ufn5+jo6Onp6erq6uvr6+zs7O3t7e7u7u/v7/Dw8PHx8fLy8vPz8/T09PX19fb29vf39/j4+Pn5+fr6+vv7+/z8/P39/f7+/v///yH5BAAAAAAAIf4fR2VuZXJhdGVkIGJ5IG9ubGluZUdJRnRvb2xzLmNvbQAsAAAAAAEAAQAACAQAAQQEADs=",
			wantError: false,
		},
		{
			name:      "Invalid image",
			image:     "base64image2",
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var (
				client = openai.NewClient(key)
				req    = openai.NewRequest("English", "", tt.image, []string{"keyword1", "keyword2"})
			)

			_, err := client.Do(context.Background(), req)
			if tt.wantError && err == nil {
				t.Error("Expected an error, got none")
			}

			if !tt.wantError && err != nil {
				t.Errorf("Did not expect an error, but got: %v", err)
			}
		})
	}
}
