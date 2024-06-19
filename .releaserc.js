module.exports = {
  "branches": [
    "main",
    {
      "name": "test",
      "prerelease": true
    }
  ],
  "plugins": [
    [
      '@semantic-release/commit-analyzer',
      {
        'config': '@fingerprintjs/conventional-changelog-dx-team',
        'releaseRules': '@fingerprintjs/conventional-changelog-dx-team/release-rules'
      }
    ],
    [
      "@semantic-release/release-notes-generator",
      {
        'config': '@fingerprintjs/conventional-changelog-dx-team',
      }
    ],
    [
      "@semantic-release/changelog",
      {
        "changelogFile": "CHANGELOG.md"
      }
    ],
    [
      "@semantic-release/exec",
      {
        "prepareCmd": "VERSION=\"${nextRelease.version}\" go run generate.go"
      }
    ],
    [
      "@semantic-release/git",
      {
        "assets": [
          "CHANGELOG.md",
          "README.md",
          "config.json",
          "**/*.go",
          "go.mod",
          "template/README.mustache"
        ]
      }
    ],
    "@semantic-release/github"
  ]
}
