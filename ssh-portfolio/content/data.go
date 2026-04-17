package content

type Profile struct {
	Name     string
	Role     string
	Location string
	Email    string
	Bio      []string
}

type Job struct {
	Title   string
	Company string
	Period  string
	Bullets []string
	Tags    []string
}

type Project struct {
	Name        string
	Event       string
	Description string
	Tags        []string
	URL         string
}

type SkillGroup struct {
	Category string
	Items    []string
}

type Link struct {
	Label string
	URL   string
}

var MyProfile = Profile{
	Name:     "Samardeep Singh",
	Role:     "Engineering Student @ University of Alberta",
	Location: "Edmonton, AB, Canada",
	Email:    "samarde1@ualberta.ca",
	Bio: []string{
		"Engineering student at the University of Alberta, transferring into the BSc Computing Science – AI Major (Fall 2026). Focused on reinforcement learning, graph neural networks, and computational drug discovery.",
		"Previously 15 months as a Data Analyst / ML Intern at Molecule AI — drug-target interaction analysis, molecular toxicity classifiers, SMILES preprocessing pipelines.",
		"Software team member at ARVP building ROS2 computer vision pipelines and the team's first RL integration for autonomous underwater navigation on Kenai, an AUV competing in RoboSub.",
	},
}

var Jobs = []Job{
	{
		Title:   "Data Analyst / Machine Learning Intern",
		Company: "Molecule AI Pvt. Ltd. — New Delhi, India",
		Period:  "Mar 2023 – Jun 2024",
		Bullets: []string{
			"Engineered end-to-end preprocessing pipelines for molecular property datasets (SMILES strings, bioactivity labels) enabling reproducible feature engineering for GNN and Ridge Regression models.",
			"Conducted EDA on drug-target interaction datasets; statistical modeling identified physicochemical feature correlations that improved virtual screening hit rates.",
			"Trained and benchmarked scikit-learn classifiers for molecular toxicity prediction, reducing false-positive rates in compound filtering and accelerating lead prioritization workflows.",
		},
		Tags: []string{"PyTorch", "GNN", "Scikit-learn", "Python", "pandas", "RDKit", "Drug Discovery"},
	},
	{
		Title:   "Asset Protection Associate",
		Company: "Walmart Canada — Edmonton, AB",
		Period:  "Oct 2025 – Present",
		Bullets: []string{
			"Execute systematic inventory audits and data-driven loss investigations using internal reporting tools.",
			"Produce structured incident reports and root-cause analyses that identify anomalous patterns and inform process improvements.",
			"Coordinate with department managers to implement corrective action plans reducing documented inventory discrepancies.",
		},
		Tags: []string{"Data Analysis", "Reporting", "Operations"},
	},
	{
		Title:   "Software Team Member – Perception & RL Integration",
		Company: "Autonomous Robotic Vehicle Project (ARVP), University of Alberta",
		Period:  "Oct 2025 – Present",
		Bullets: []string{
			"Developing ROS2-based computer vision pipelines for Kenai, ARVP's AUV competing in the international RoboSub Competition; implementing YOLO-based underwater object detection integrated into mission-planning behaviour trees.",
			"Pioneering the team's first RL integration — designing a reward-structured simulation environment to replace hand-tuned controllers with a deep RL policy for autonomous underwater navigation.",
		},
		Tags: []string{"ROS2", "Python", "YOLO", "OpenCV", "Reinforcement Learning", "Behaviour Trees"},
	},
}

var Projects = []Project{
	{
		Name:        "TorsionNet — Deep RL for Molecular Conformer Generation",
		Event:       "HackED 2026",
		Description: "Replicated TorsionNet (NeurIPS), a deep RL approach to conformer search. Built a custom OpenAI Gym environment with RDKit, trained a PPO agent with an MPNN backbone and curriculum learning, deployed inference via FastAPI and Docker. Achieved 90% reduction in generation time vs. brute-force search.",
		Tags:        []string{"PyTorch", "PyTorch Geometric", "RDKit", "PPO", "GNN", "FastAPI", "Docker"},
		URL:         "",
	},
	{
		Name:        "CoSound — Mindful Social Listening",
		Event:       "NatHacks 2024 · Honorary Mention",
		Description: "Real-time ambient-audio-driven music recommendation system, physically installed at Cameron Library, University of Alberta. Ridge Regression classifier on ESC-50 dataset; 5-dimensional user preference embeddings from Librosa audio features matched against song profiles via live NFC-tag voting.",
		Tags:        []string{"Python", "Scikit-learn", "Librosa", "React", "Node.js", "PostgreSQL", "Docker"},
		URL:         "",
	},
	{
		Name:        "GraphSAGE-MAPPO",
		Event:       "",
		Description: "Multi-agent RL system on mesh networks using GraphSAGE and MAPPO. Built for research positioning at the intersection of GNNs and multi-agent systems.",
		Tags:        []string{"PyTorch", "GNN", "MARL", "Python"},
		URL:         "",
	},
	{
		Name:        "DrugBAN Reimplementation (In Progress)",
		Event:       "",
		Description: "Reimplementing the Bilinear Attention Network for drug-target interaction prediction. Next project in computational drug discovery portfolio following TorsionNet.",
		Tags:        []string{"PyTorch", "Bioinformatics", "Attention", "GNN"},
		URL:         "",
	},
}

var SkillGroups = []SkillGroup{
	{
		Category: "ML / AI",
		Items:    []string{"PyTorch", "PyTorch Geometric", "Scikit-learn", "Deep Learning", "PPO / RL", "GNN", "Statistical Modeling", "NumPy", "SciPy", "Librosa"},
	},
	{
		Category: "Data Science",
		Items:    []string{"Python", "pandas", "SQL", "PostgreSQL", "Feature Engineering", "EDA", "Jupyter", "matplotlib", "seaborn"},
	},
	{
		Category: "Robotics / CV",
		Items:    []string{"ROS2", "OpenCV", "YOLO", "Behaviour Trees", "AUV Stack", "3D CAD"},
	},
	{
		Category: "Engineering / MLOps",
		Items:    []string{"Docker", "Git", "SSH", "Linux", "FastAPI", "REST APIs", "CI/CD", "AWS", "Supabase", "React", "GoLang"},
	},
}

var Links = []Link{
	{Label: "GitHub", URL: "https://github.com/MoveFastAndBreakThings-dot?tab=repositories"},
	{Label: "LinkedIn", URL: "https://www.linkedin.com/in/samardeep-singh-488b4024b/"},
	{Label: "Email", URL: "mailto:samarde1@ualberta.ca"},
	{Label: "Resume", URL: "https://drive.google.com/file/d/1bvnNffGLOVFWwAoVa1sZzd6BVQ-xCNEX/view?usp=sharing"},
}

var Extracurriculars = []string{
	"Runner-Up — Enbridge Case Competition, University of Alberta (2025)",
	"Runner-Up — Global Health Case Competition, MSF / Hep-C in Cox's Bazar (2026)",
	"Participant — HackED 2026, University of Alberta Engineering Students' Society",
	"Honorary Mention — NatHacks 2024, NeurAlbertaTech Hacks",
}
