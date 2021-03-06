package slugify

import "testing"

func TestSlugify(t *testing.T) {

	results := make(map[string]string)
	results["hello-playground"] = "Hello, playground"
	results["hello-it-s-paradise"] = "😢 😣 😤 😥 😦 😧 😨 😩 😪 😫 😬 Hello, it's paradise"
	results["hi-this-is-a-test"] = "方向盤後面 hi this is a test خلف المقو"

	for slug, original := range results {
		actual := Slugify(original)

		if actual != slug {
			t.Errorf("Expected '%s', got: %s", slug, actual)
		}
	}

}

func TestCustomSlugifier(t *testing.T) {

	slugifier := New(Configuration{ReplaceCharacter: '.'})

	results := make(map[string]string)
	results["hello.playground"] = "Hello, playground"
	results["hello.it.s.paradise"] = "😢 😣 😤 😥 😦 😧 😨 😩 😪 😫 😬 Hello, it's paradise"
	results["hi.this.is.a.test"] = "方向盤後面 hi this is a test خلف المقو"

	for slug, original := range results {
		actual := slugifier.Slugify(original)

		if actual != slug {
			t.Errorf("Expected '%s', got: %s", slug, actual)
		}
	}

}

func TestCustomSlugifierWithChecker(t *testing.T) {

	slugifier := New(Configuration{
		IsValidCharacterChecker: func(c uint8) bool {
			if c >= 'a' && c <= 'z' {
				return true
			}

			return false
		},
	})

	results := make(map[string]string)
	results["hello-playground"] = "Hello, playground"
	results["hello-it-s-paradise"] = "Hello, it's 123 paradise"
	results["hi-i-s-a-test"] = "hi 091 i3s a test"

	for slug, original := range results {
		actual := slugifier.Slugify(original)

		if actual != slug {
			t.Errorf("Expected '%s', got: %s", slug, actual)
		}
	}
}

func BenchmarkSlugify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slugify("Hello, world!")
	}
}

func BenchmarkSlugifyLongString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slugify(`
			😢 😣 😤 😥 😦 😧 😨 😩 😪 😫 😬 Hello, it's paradise
			😢 😣 😤 😥 😦 😧 😨 😩 😪 😫 😬 Hello, it's paradise
			😢 😣 😤 😥 😦 😧 😨 😩 😪 😫 😬 Hello, it's paradise
			😢 😣 😤 😥 😦 😧 😨 😩 😪 😫 😬 Hello, it's paradise
			Lorem ipsum dolor sit amet, consectetur adipiscing elit.
			Aliquam sapien nisl, laoreet quis vestibulum ut, cursus 
			in turpis. Sed magna mi, blandit id nisi vel, imperdiet 
			mollis turpis. Fusce vel fringilla mauris. Donec cursus 
			rhoncus bibendum. Aliquam erat volutpat. Maecenas 
			faucibus turpis ex, quis lacinia ligula ultrices non. 
			Sed gravida justo augue. Nulla bibendum dignissim tellus 
			vitae lobortis. Suspendisse fermentum vel purus in pulvinar. 
			Vivamus eu fermentum purus, sit amet tempor orci. 
			Praesent congue convallis turpis, ac ullamcorper lorem 
			semper id. 
		`)
	}
}
