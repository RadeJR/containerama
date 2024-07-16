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
	import { v4 as uuidv4 } from "uuid";
	import Spinner from "$lib/components/ui/spinner/spinner.svelte";
	import * as Tabs from "$lib/components/ui/tabs";
	import * as Card from "$lib/components/ui/card/index.js";
	import { title } from "$store";
	import { toast } from "svelte-sonner";
	import axios from "axios";

	title.set("Create a Stack");

	var isLoading = false;
	let createWebhook: boolean = false;
	let webhookUrl: string;

	$: if (createWebhook == true) {
		data.webhook = uuidv4();
		webhookUrl =
			window.location.protocol +
			"//" +
			window.location.hostname +
			":" +
			window.location.port +
			"/webhook/" +
			data.webhook;
	} else {
		data.webhook = "";
	}

	let data: StackFileData = new StackFileData();
	async function send() {
		isLoading = true;
		await axios
			.post("/api/stacks", data)
			.then(() => {
				push("/stacks");
				toast.success("Stack has been created");
			})
			.catch(function (error) {
				if (error.response) {
					toast.error("An error occured", {
						description: error.response.data.message.join("\n"),
					});
				} else if (error.request) {
					console.log(error.request);
				} else {
					// Something happened in setting up the request that triggered an Error
					console.log("Error", error.message);
				}
			})
			.finally(() => {
				isLoading = false;
			});
	}
</script>

<div transition:fade={{ duration: 100 }}>
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
			<div class="flex justify-between items-center gap-2 h-9 my-5">
				<div class="flex items-center gap-3">
					<Switch id="createwebhook" bind:checked={createWebhook} />
					<Label for="createwebhook">Create Webhook</Label>
				</div>
				{#if createWebhook}
					<div class="flex items-center gap-3">
						<Label>{webhookUrl}</Label>
						<Button
							on:click={() =>
								navigator.clipboard.writeText(webhookUrl)}
							>Copy</Button
						>
					</div>
				{/if}
			</div>
		</div>
		<Tabs.Root value="git" class="w-full">
			<Tabs.List class="h-20 grid w-full grid-cols-2">
				<Tabs.Trigger class="h-16 text-xl" value="git"
					>Git Repo</Tabs.Trigger
				>
				<Tabs.Trigger class="h-16 text-xl" value="file"
					>Compose File</Tabs.Trigger
				>
			</Tabs.List>
			<Tabs.Content value="git">
				<div class="w-full h-[200px] content-center items-center flex justify-around">
					<h1>Comming Soon</h1>
				</div>
			</Tabs.Content>
			<Tabs.Content value="file">
				<Card.Root>
					<Card.Header>
						<Card.Title>File</Card.Title>
						<Card.Description>
							Enter your compose file here. Your stack will be
							deployed using "docker compose up".
						</Card.Description>
					</Card.Header>
					<Card.Content class="space-y-2">
						<YamlEditor bind:content={data.content} />
					</Card.Content>
					<Card.Footer>
						<Button
							on:click={send}
							disabled={isLoading}
							class="w-full"
						>
							{#if isLoading}
								<Spinner />
							{/if}
							Create</Button
						>
					</Card.Footer>
				</Card.Root>
			</Tabs.Content>
		</Tabs.Root>
	</div>
</div>
