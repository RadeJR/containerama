interface PortBinding {
	HostIp: string;
	HostPort: string;
}

interface Mount {
	Type: string;
	Source: string;
	Target: string;
	VolumeOptions?: object;
}

interface PortBindings {
	[key: string]: PortBinding[];
}

interface Labels {
	[key: string]: string;
}

function portBindingsToString(portBindings: PortBindings): string {
	const result: string[] = [];

	for (const port in portBindings) {
		if (portBindings.hasOwnProperty(port)) {
			const bindings = portBindings[port];
			bindings.forEach(binding => {
				const hostIp = binding.HostIp || "0.0.0.0"; // Default to 0.0.0.0 if HostIp is empty
				result.push(`${hostIp}:${binding.HostPort}:${port}`);
			});
		}
	}

	return result.join('\n');
}

function labelsToString(labels: Labels): string {
	const result: string[] = [];

	for (const key in labels) {
		if (labels.hasOwnProperty(key)) {
			result.push(`${key}=${labels[key]}`);
		}
	}

	return result.join('\n');
}


function mountsToString(mounts: Mount[]): string {
	const result: string[] = mounts.map(mount => `${mount.Source}:${mount.Target}`);
	return result.join('\n');
}

export { portBindingsToString, labelsToString, mountsToString };
