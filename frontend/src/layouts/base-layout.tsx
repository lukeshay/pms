import { FC, ReactNode } from "react";
import { AuthClaims } from "../api";
import { Link, useNavigate } from "react-router-dom";
import { useTokens } from "../hooks/use-tokens";
import { XIcon } from "lucide-react";
import { Head } from "../components/head";

export type BaseLayoutProps = {
	children: ReactNode;
	claims?: AuthClaims;
	splitPanel?: ReactNode;
	onSplitPanelClose?: () => void;
	title?: string;
};

export const BaseLayout: FC<BaseLayoutProps> = ({ children, claims, splitPanel, onSplitPanelClose, title = "PMS" }) => {
	const { setTokens } = useTokens();
	const navigate = useNavigate();

	return (
		<>
			<Head>
				<title>{title}</title>
			</Head>
			<div className="flex min-h-screen bg-base-200">
				<div className="flex w-full justify-center px-3 py-6">
					<div className="w-full max-w-6xl">
						<nav className="navbar rounded-box bg-base-100">
							<div className="flex-1">
								<Link to="/" className="btn-ghost btn text-xl normal-case">
									daisyUI
								</Link>
							</div>
							<div className="flex-none">
								<ul className="menu menu-horizontal px-1">
									{claims && (
										<li>
											<Link to="/">{"Books"}</Link>
										</li>
									)}
									<li>
										<details>
											<summary>Profile</summary>
											<ul className="bg-base-100 p-2">
												{claims ? (
													<>
														<li>
															<Link to="/profile">{"Profile"}</Link>
														</li>
														<li>
															<button
																onClick={() => {
																	setTokens();
																	navigate("/sign-in");
																}}
															>
																{"Sign Out"}
															</button>
														</li>
													</>
												) : (
													<li>
														<Link to="/sign-in">{"Sign In"}</Link>
													</li>
												)}
											</ul>
										</details>
									</li>
								</ul>
							</div>
						</nav>
						<section id="content" className="pb-16 pt-8">
							{children}
						</section>
					</div>
				</div>
				{splitPanel && (
					<div className="min-w-[30rem] border border-base-100">
						<div>
							<button
								onClick={() => {
									onSplitPanelClose();
								}}
								className="btn-ghost btn-sm btn-circle float-right flex items-center justify-center"
							>
								<XIcon />
							</button>
						</div>
						<div className="p-4 pt-8">{splitPanel}</div>
					</div>
				)}
			</div>
		</>
	);
};
