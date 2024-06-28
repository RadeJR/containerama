<script lang="ts">
	import { fade } from "svelte/transition";
	import { Input } from "$lib/components/ui/input/index.js";
	import { Label } from "$lib/components/ui/label";
	import { Button } from "$lib/components/ui/button";
	import { Switch } from "$lib/components/ui/switch/index.js";
	import { StackFileData } from "$app/types/stack";
	import { getAxios } from "$conf/axios";
	import { push } from "svelte-spa-router";
	import YamlEditor from "./YamlEditor.svelte";
	import {v4 as uuidv4} from 'uuid';

	let useCompose: boolean = false;
	let createWebhook: boolean = false;

	$: if (createWebhook == true) {
		data.webhook = uuidv4();
	} else {
		data.webhook = "";
	}

	let data: StackFileData = new StackFileData();
	async function send() {
		await getAxios()
			.post("/api/stacks", data)
			.then(() => {
				push("/stacks");
			});
	}
</script>

<div transition:fade={{ duration: 100 }}>
	<div class="flex justify-between py-4">
		<h4 class="text-lg font-medium">Create Stack</h4>
		<div class="flex gap-3">
			<div class="flex items-center gap-2">
				<Switch id="usecompose" bind:checked={useCompose} />
				<Label for="usecompose">Use compose file</Label>
			</div>
			<Button on:click={send} variant="outline">Create</Button>
		</div>
	</div>
	<div class="rounded-md border p-5">
		<div class="flex flex-col my-1 gap-2">
			<div class="flex-1">
				<Label for="name">Stack Name</Label>
				<Input
					id="name"
					type="text"
					placeholder="name"
					bind:value={data.name}
				/>
			</div>
			<div class="flex items-center gap-2">
				<Switch id="createwebhook" bind:checked={createWebhook} />
				<Label for="createwebhook">Create Webhook</Label>
			</div>
		</div>
		{#if useCompose}
			<div class="flex-1">
				<Label for="image">Content</Label>
				<YamlEditor bind:content={data.content} />
			</div>
		{/if}
	</div>
</div>
