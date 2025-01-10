package flags

var Structure = map[string]interface{}{
	"cmd": map[string]interface{}{
		"api": map[string]interface{}{
			"main.go": "template",
		},
		"web": map[string]interface{}{},
	},
	"internal": map[string]interface{}{
		"config": map[string]interface{}{
			"app.go":       "template",
			"config.go":    "template",
			"fiber.go":     "template",
			"database.go":  "template",
			"logrus.go":    "template",
			"validator.go": "template",
		},
		"domain": map[string]interface{}{
			"User.go": "template",
			"Role.go": "template",
		},
		"repository": map[string]interface{}{
			"user.repository.go": "template",
			"role.repository.go": "template",
		},
		"service": map[string]interface{}{
			"user.service.go": "template",
			"role.service.go": "template",
		},
		"dto": map[string]interface{}{
			// "user.dto.go": nil,
			// "role.dto.go": nil,
		},
		"routes": map[string]interface{}{
			"handler": map[string]interface{}{
				// "user.handler.go": nil,
				// "role.handler.go": nil,
			},
			// "response.go": nil,
			// "routes.go":   nil,
		},
	},
	"pkg": map[string]interface{}{
		"middleware": map[string]interface{}{
			"auth.go": "template",
		},
	},
	"utils":        map[string]interface{}{},
	".env":         nil,
	".env.example": nil,
	".gitignore":   nil,
	"Readme.md":    "# Project Title\n\nThis is a generated project structure.\n",
}
