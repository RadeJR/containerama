<script lang="ts">
	import { fade } from "svelte/transition";
	import { Input } from "$lib/components/ui/input/index.js";
	import { Label } from "$lib/components/ui/label";
	import { Button } from "$lib/components/ui/button";
	import { Textarea } from "$lib/components/ui/textarea/index.js";
	import { Switch } from "$lib/components/ui/switch/index.js";
	import { ContainerData } from "$app/types/container";
	import { getAxios } from "$conf/axios";
	import { push } from "svelte-spa-router";

	let data: ContainerData = new ContainerData();
	async function send() {
		await getAxios()
			.post("/api/containers/create", data)
			.then(() => {
				push("/containers");
			});
	}
</script>

<div transition:fade={{ duration: 100 }}>
	<div class="flex justify-between py-4">
		<h4 class="text-lg font-medium">Create Container</h4>
		<Button on:click={send} variant="outline">Create</Button>
	</div>
	<div class="rounded-md border p-5">
		<div class="flex my-1 gap-2">
			<div class="flex-1">
				<Label for="name">Container Name</Label>
				<Input
					id="name"
					type="text"
					placeholder="name"
					bind:value={data.name}
				/>
			</div>
			<div class="flex-1">
				<Label for="image">Container image</Label>
				<Input
					id="image"
					type="text"
					placeholder="image"
					bind:value={data.image}
				/>
			</div>
		</div>
		<div class="flex my-1 gap-2">
			<div class="flex-1">
				<Label for="cmd">Container command</Label>
				<Input
					id="cmd"
					type="text"
					placeholder="command"
					bind:value={data.cmd}
				/>
			</div>
			<div class="flex-1">
				<Label for="entrypoint">Container entrypoint</Label>
				<Input
					id="entrypoint"
					type="text"
					placeholder="entrypoint"
					bind:value={data.entrypoint}
				/>
			</div>
		</div>
		<div>
			<Switch id="networking" bind:checked={data.networkDisabled} />
			<Label for="networking">Networking</Label>
		</div>
		<div>
			<Label for="volumes">Container volumes</Label>
			<Textarea
				id="volumes"
				placeholder="volume:/path"
				bind:value={data.volumes}
			/>
		</div>
		<div class="my-1">
			<Label for="env">Environment variables</Label>
			<Textarea
				id="env"
				placeholder="APP_ENV=...."
				bind:value={data.env}
			/>
		</div>
		<div class="my-1">
			<Label for="ports">Container ports</Label>
			<Textarea
				id="ports"
				placeholder="8080:80"
				bind:value={data.ports}
			/>
		</div>
		<div class="my-1">
			<Label for="labels">Container labels</Label>
			<Textarea
				id="labels"
				placeholder="Labels"
				bind:value={data.labels}
			/>
		</div>
	</div>
</div>
