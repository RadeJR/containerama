<script lang="ts">
	import { onMount } from "svelte";
	import { EditorView, basicSetup } from "codemirror";
	import { EditorState } from "@codemirror/state";
	import { yaml } from "@codemirror/lang-yaml";
	import { keymap } from "@codemirror/view";
	import { indentWithTab } from "@codemirror/commands";
	import { catppuccin } from "codemirror-theme-catppuccin";

	const variant = "mocha";
	const theme = catppuccin(variant);
	const mytheme = EditorView.theme({
		".cm-gutter,.cm-content": { minHeight: "500px" },
		".cm-scroller": { overflow: "auto" },
	});
	let editorDiv: HTMLElement;
	let view: EditorView;
	export let content = "";

	let updateListenerExtension = EditorView.updateListener.of((update) => {
		if (update.docChanged) {
			content = view.state.doc.toString();
		}
	});

	onMount(() => {
		const state = EditorState.create({
			doc: "",
			extensions: [
				basicSetup,
				yaml(),
				keymap.of([indentWithTab]),
				theme,
				mytheme,
				updateListenerExtension,
			],
		});

		view = new EditorView({
			state,
			parent: editorDiv,
		});

		return () => {
			view.destroy();
		};
	});
</script>

<div bind:this={editorDiv}></div>

<style>
	div {
		font-family: monospace;
		width: 100%;
	}
</style>
