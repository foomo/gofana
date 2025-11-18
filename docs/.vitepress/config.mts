import { defineConfig } from 'vitepress'
import {
	groupIconMdPlugin,
	groupIconVitePlugin,
} from 'vitepress-plugin-group-icons'
import llmstxt, { copyOrDownloadAsMarkdownButtons } from 'vitepress-plugin-llms'

const version = 'latest'

export default defineConfig({
	title: 'gofana',
	description: 'CLI utility to manage Grafana resources',
	themeConfig: {
		logo: '/logo.png',
		nav: [
			{
				text: `${version}`,
				items: [
					{
						text: 'Release Notes',
						link: 'https://github.com/foomo/gofana/releases',
					},
				],
			},
		],
		sidebar: [
			{
				text: 'Guide',
				items: [
					{ text: 'Introduction', link: '/' },
				],
			},
			{
				text: 'Reference',
				items: [
					{
						text: 'Cli',
						link: 'reference/cli/gofana.md',
						items: [
							{ text: 'config', link: 'reference/cli/gofana_config.md' },
							{ text: 'generate', link: 'reference/cli/gofana_generate.md' },
							{ text: 'list', link: 'reference/cli/gofana_list.md' },
							{ text: 'version', link: 'reference/cli/gofana_version.md' },
						],
					},
				],
			},
			{
				text: 'Contributing',
				items: [
					{
						text: "Guideline",
						link: '/CONTRIBUTING.md',
					},
					{
						text: "Code of conduct",
						link: '/CODE_OF_CONDUCT.md',
					},
					{
						text: "Security guidelines",
						link: '/SECURITY.md',
					},
				],
			},
		],
		editLink: {
			pattern: 'https://github.com/foomo/gofana/edit/main/docs/:path',
			text: 'Suggest changes to this page',
		},
		search: {
			provider: 'local',
		},
		footer: {
			message: 'Released under the MIT License.',
		},
		socialLinks: [
			{
				icon: 'github',
				link: 'https://github.com/foomo/gofana',
			},
		],
	},
	head: [
		['meta', { name: 'theme-color', content: '#ffffff' }],
		['link', { rel: 'icon', href: '/logo.png' }],
		['meta', { name: 'author', content: 'foomo by bestbytes' }],
		['meta', { property: 'og:title', content: 'foomo/gofana' }],
		[
			'meta',
			{
				property: 'og:image',
				content: 'https://github.com/foomo/gofana/blob/main/docs/public/banner.png?raw=true',
			},
		],
		[
			'meta',
			{
				property: 'og:description',
				content: 'CLI utility to manage Grafana resources',
			},
		],
		['meta', { name: 'twitter:card', content: 'summary_large_image' }],
		[
			'meta',
			{
				name: 'twitter:image',
				content: 'https://github.com/foomo/gofana/blob/main/docs/public/banner.png?raw=true',
			},
		],
		[
			'meta',
			{
				name: 'viewport',
				content: 'width=device-width, initial-scale=1.0, viewport-fit=cover',
			},
		],
	],
	markdown: {
		theme: {
			dark: 'one-dark-pro',
			light: 'github-light',
		},
		config(md) {
			md.use(groupIconMdPlugin)
			md.use(copyOrDownloadAsMarkdownButtons)
		},
	},
	vite: {
		plugins: [
			groupIconVitePlugin(),
			llmstxt({
				excludeIndexPage: false,
			}),
		],
	},
	sitemap: {
		hostname: 'https://foomo.github.io/gofana',
	},
	ignoreDeadLinks: true,
})
