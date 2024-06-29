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
	import { onMount } from "svelte";
	import {
		portBindingsToString,
		labelsToString,
		mountsToString,
	} from "./util";

	let data: ContainerData = new ContainerData();
	export let params: any = {};

	async function get() {
		await getAxios()
			.get(`/api/containers/${params.id}`)
			.then((resp) => {
				data.name = resp.data.Name;
				data.labels = labelsToString(resp.data.Config.Labels);
				data.ports = portBindingsToString(
					resp.data.HostConfig.PortBindings,
				);
				data.env = resp.data.Config.Env.join("\n");
				data.image = resp.data.Config.Image;
				data.cmd = resp.data.Config.Cmd.join("\n");
				data.volumes = mountsToString(resp.data.HostConfig.Mounts);
				data.entrypoint = resp.data.Config.Entrypoint.join("\n");
				console.log(data);
			});
	}

	onMount(() => {
		get();
	});

	async function send() {
		await getAxios()
			.patch(`/api/containers/${params.id}`, data)
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
			<Label for="networking">Disable networking</Label>
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
			<Textarea id="env" placeholder="key=value" bind:value={data.env} />
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
				placeholder="key=value"
				bind:value={data.labels}
			/>
		</div>
	</div>
</div>
